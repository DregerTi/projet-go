package adapter

import (
	"context"
	"errors"
	pb "finalProject/gen/proto/go"
	"finalProject/helpers"
	"finalProject/payment"
	"finalProject/product"
	"finalProject/service"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
	roomManager service.Manager
}

func NewTestApiServer(rm service.Manager) pb.TestApiServer {
	return &testApiServer{roomManager: rm}
}

func (s *testApiServer) GetProductById(ctx context.Context, request *pb.ProductByIdRequest) (*pb.ProductByIdResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}

	var productToFind product.Product
	//TODO change id ty
	id := request.GetId()
	err = helpers.TestIdValidity(request.GetId())
	if err != nil {
		return nil, err
	}

	db.First(&productToFind, id)

	return &pb.ProductByIdResponse{
		Id:        productToFind.Id,
		Name:      productToFind.Name,
		Price:     productToFind.Price,
		UpdatedAt: timestamppb.New(productToFind.UpdatedAt),
		CreatedAt: timestamppb.New(productToFind.CreatedAt),
	}, nil
}

func (s *testApiServer) CreateProduct(c context.Context, request *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}

	actualTime := time.Now()
	price := request.GetPrice()
	err = helpers.TestPriceValidity(price)
	if err != nil {
		return nil, err
	}
	name := request.GetName()
	productToCreate := &product.Product{
		Name:      name,
		Price:     price,
		UpdatedAt: actualTime,
		CreatedAt: actualTime,
	}

	if db.First(&productToCreate, "name = ?", name).RowsAffected != 0 {
		return nil, errors.New("product already exists with name " + name)
	}
	db.Create(&productToCreate)

	return &pb.CreateProductResponse{
		Name:      request.GetName(),
		Price:     request.GetPrice(),
		UpdatedAt: timestamppb.New(actualTime),
		CreatedAt: timestamppb.New(actualTime),
	}, nil
}

func (s *testApiServer) UpdateProduct(c context.Context, request *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}

	id := request.GetId()
	err = helpers.TestIdValidity(id)
	if err != nil {
		return nil, err
	}
	var productToUpdate product.Product
	if db.First(&productToUpdate, id).Error != nil {
		errorNotFound := fmt.Errorf("product with id %d not found", id)
		return nil, errors.New(errorNotFound.Error())
	}
	name := request.GetName()
	if name != "" {
		if db.First(&productToUpdate, "name = ?", name).RowsAffected != 0 {
			return nil, errors.New("product already exists with name " + name)
		}
		productToUpdate.Name = request.GetName()
	}
	price := request.GetPrice()
	if price != 0 {
		err = helpers.TestPriceValidity(price)
		if err != nil {
			return nil, err
		}
		productToUpdate.Price = price
	}
	productToUpdate.UpdatedAt = time.Now()
	db.Save(&productToUpdate)

	return &pb.UpdateProductResponse{
		Id:        productToUpdate.Id,
		Name:      productToUpdate.Name,
		Price:     productToUpdate.Price,
		UpdatedAt: timestamppb.New(productToUpdate.UpdatedAt),
		CreatedAt: timestamppb.New(productToUpdate.CreatedAt),
	}, nil
}

func (s *testApiServer) GetAllProducts(c context.Context, request *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}

	var products []product.Product
	err = db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	productsResponse := make([]*pb.ProductByIdResponse, len(products))
	for i, product := range products {
		productsResponse[i] = &pb.ProductByIdResponse{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			UpdatedAt: timestamppb.New(product.UpdatedAt),
			CreatedAt: timestamppb.New(product.CreatedAt),
		}
	}
	return &pb.GetAllProductsResponse{
		Products: productsResponse,
	}, nil
}

func (s *testApiServer) DeleteProductById(ctx context.Context, request *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}

	id := request.GetId()
	err = helpers.TestIdValidity(id)
	if err != nil {
		return nil, err
	}
	productToDelete := &product.Product{Id: id}
	errorDelete := db.Delete(productToDelete)

	if errorDelete.Error != nil {
		return nil, errorDelete.Error
	}

	if errorDelete.RowsAffected == 0 {
		errorNotFound := fmt.Errorf("product with id %d not found", id)
		return nil, errors.New(errorNotFound.Error())
	}
	return &pb.DeleteProductByIdResponse{
		Message: "Product deleted",
	}, nil
}

func (s *testApiServer) CreatePayment(c context.Context, request *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}

	productBought := &product.Product{Id: request.GetProductId()}
	err = db.First(productBought).Error
	if err != nil {
		return nil, err
	}

	actualTime := time.Now()
	paymentToCreate := &payment.Payment{
		ProductId: request.GetProductId(),
		PricePaid: productBought.Price,
		UpdatedAt: actualTime,
		CreatedAt: actualTime,
	}

	err = db.Create(&paymentToCreate).Error
	if err != nil {
		return nil, err
	}

	_, err = s.PostToRoom(c, &pb.Message{
		UserId: "1",
		RoomId: "payment",
		Text:   helpers.PaymentToString(paymentToCreate, productBought),
	})
	if err != nil {
		return nil, err
	}

	paymentToDisplay := &pb.CreatePaymentResponse{
		ProductId: request.GetProductId(),
		PricePaid: productBought.Price,
		UpdatedAt: timestamppb.New(actualTime),
		CreatedAt: timestamppb.New(actualTime),
	}

	return paymentToDisplay, nil
}

func (s *testApiServer) GetPaymentById(c context.Context, request *pb.GetPaymentByIdRequest) (*pb.GetPaymentByIdResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}
	var paymentToFind payment.Payment
	id := request.GetId()
	err = helpers.TestIdValidity(id)
	if err != nil {
		return nil, err
	}

	err = db.First(&paymentToFind, id).Error
	if err != nil {
		return nil, err
	}
	if &paymentToFind == nil {
		panic("Payment not found")
	}

	return &pb.GetPaymentByIdResponse{
		Id:        paymentToFind.Id,
		ProductId: paymentToFind.ProductId,
		PricePaid: paymentToFind.PricePaid,
		UpdatedAt: timestamppb.New(paymentToFind.UpdatedAt),
		CreatedAt: timestamppb.New(paymentToFind.CreatedAt),
	}, nil
}

func (s *testApiServer) UpdatePayment(c context.Context, request *pb.UpdatePaymentRequest) (*pb.UpdatePaymentResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}
	var paymentToUpdate payment.Payment
	id := request.GetId()
	if err != nil {
		return nil, err
	}

	err = helpers.TestIdValidity(id)
	if err != nil {
		return nil, err
	}
	db.First(&paymentToUpdate, id)
	if &paymentToUpdate == nil {
		return nil, errors.New("payment not found")
	}

	paymentToUpdate.ProductId = request.GetProductId()
	paymentToUpdate.PricePaid = request.GetPricePaid()
	paymentToUpdate.UpdatedAt = time.Now()
	err = db.Save(&paymentToUpdate).Error
	if err != nil {
		return nil, err
	}

	productBought := &product.Product{Id: request.GetProductId()}
	err = db.First(productBought).Error
	if err != nil {
		return nil, err
	}

	_, err = s.PostToRoom(c, &pb.Message{
		UserId: "1",
		RoomId: "payment",
		Text:   helpers.PaymentToString(&paymentToUpdate, productBought),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePaymentResponse{
		Id:        paymentToUpdate.Id,
		ProductId: paymentToUpdate.ProductId,
		PricePaid: paymentToUpdate.PricePaid,
		UpdatedAt: timestamppb.New(paymentToUpdate.UpdatedAt),
		CreatedAt: timestamppb.New(paymentToUpdate.CreatedAt),
	}, nil
}

func (s *testApiServer) GetAllPayments(c context.Context, request *pb.GetAllPaymentsRequest) (*pb.GetAllPaymentsResponse, error) {
	db, err := helpers.GetDatabase()
	if err != nil {
		return nil, err
	}
	var payments []payment.Payment
	err = db.Find(&payments).Error
	if err != nil {
		return nil, err
	}

	paymentsResponse := make([]*pb.GetPaymentByIdResponse, len(payments))
	for i, payment := range payments {
		paymentsResponse[i] = &pb.GetPaymentByIdResponse{
			Id:        payment.Id,
			ProductId: payment.ProductId,
			PricePaid: payment.PricePaid,
			UpdatedAt: timestamppb.New(payment.UpdatedAt),
			CreatedAt: timestamppb.New(payment.CreatedAt),
		}
	}

	return &pb.GetAllPaymentsResponse{
		Payments: paymentsResponse,
	}, nil
}

func (s *testApiServer) GetRoom(ctx context.Context, rq *pb.RoomRequest) (*pb.RoomResponse, error) {
	return nil, nil
}

func (s *testApiServer) PostToRoom(ctx context.Context, msg *pb.Message) (*pb.RoomResponse, error) {
	s.roomManager.Submit(msg.UserId, msg.RoomId, msg.Text)
	return &pb.RoomResponse{
		Success: true,
	}, nil
}

func (s *testApiServer) DeleteRoom(ctx context.Context, rq *pb.RoomRequest) (*pb.RoomResponse, error) {
	return nil, nil
}

func (s *testApiServer) StreamPayments(rr *pb.RoomRequest, srs pb.TestApi_StreamPaymentsServer) error {
	listener := s.roomManager.OpenListener("payment")
	defer s.roomManager.CloseListener("payment", listener)

	for {
		select {
		case msg := <-listener:
			serviceMsg, ok := msg.(service.Message)
			if !ok {
				fmt.Println(msg)
				continue
			}

			if err := srs.Send(&pb.MessageResponse{
				Text: serviceMsg.Text,
			}); err != nil {
				return err
			}
		case <-srs.Context().Done():
			return nil
		}
	}
}
