package application


type productsService interface {

}

type paymentsService interface {

}

type OrdersService struct {

}

type PlaceOrderCommand struct {

}

type MarkOrderAsPaidCommand struct {

}

func NewOrdersService() {

}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {

}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}

func (s OrdersService) OrderById(id orders.Id) (orders.Order, error) {

}