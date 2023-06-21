package server

import (
	"TechnicalTestKStyleHub/domain"
	"TechnicalTestKStyleHub/src/member"
	MemberMapper "TechnicalTestKStyleHub/src/member/mappers"
	MemberRepository "TechnicalTestKStyleHub/src/member/repositories"
	MemberUseCase "TechnicalTestKStyleHub/src/member/usecases"
	"TechnicalTestKStyleHub/src/product"
	ProductMapper "TechnicalTestKStyleHub/src/product/mappers"
	ProductRepository "TechnicalTestKStyleHub/src/product/repositories"
	ProductUseCase "TechnicalTestKStyleHub/src/product/usecases"
)

type Contract struct {
	IMemberMapper      member.IMemberMapper
	IMemberRepository  member.IMemberRepository
	IMemberUseCase     member.IMemberUseCase
	IProductMapper     product.IProductMapper
	IProductRepository product.IProductRepository
	IProductUseCase    product.IProductUseCase
}

func InitContract(config domain.Config) (contract Contract) {

	// member
	contract.IMemberMapper = MemberMapper.NewMemberMapper()
	contract.IMemberRepository = MemberRepository.NewMemberRepository(config.Sql.DB)
	contract.IMemberUseCase = MemberUseCase.NewMemberUseCase(contract.IMemberMapper, contract.IMemberRepository)

	// product
	contract.IProductMapper = ProductMapper.NewProductMapper()
	contract.IProductRepository = ProductRepository.NewProductRepository(config.Sql.DB)
	contract.IProductUseCase = ProductUseCase.NewProductUseCase(contract.IProductMapper, contract.IProductRepository)

	return contract
}
