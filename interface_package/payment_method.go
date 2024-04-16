package interface_package

type PaymentMethod interface {
	Create() error
	Update() error
}
