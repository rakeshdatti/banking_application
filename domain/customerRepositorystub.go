

package domain 


type CustomerRepositoryStub struct{
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer,error){
	return s.customers,nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub{
	customers := []Customer{
		{"1","rakesh","rajam","532122","05/07/2002","1"},
		{"2","rocky","rajam","532122","01/01/2001","0"},
		
	}
	return CustomerRepositoryStub{customers: customers}
}
