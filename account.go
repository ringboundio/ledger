package ledger

type AccountScribe struct{ descriptor AccountDescriptor }

func NewAccountScribe() AccountScribe {
	return AccountScribe{}
}

func (this AccountScribe) Asset() AssetClassCategorySelector {
	this.descriptor.class = newAssetClass()
	return AssetClassCategorySelector{descriptor: this.descriptor}
}

func (this AccountScribe) Liability() LiabilityClassCategorySelector {
	this.descriptor.class = newLiabilityClass()
	return LiabilityClassCategorySelector{descriptor: this.descriptor}
}

func (this AccountScribe) Equity() EquityClassCategorySelector {
	this.descriptor.class = newEquityClass()
	return EquityClassCategorySelector{descriptor: this.descriptor}
}

func (this AccountScribe) Revenue() RevenueClassCategorySelector {
	this.descriptor.class = newRevenueClass()
	return RevenueClassCategorySelector{descriptor: this.descriptor}
}

func (this AccountScribe) Expense() ExpenseClassCategorySelector {
	this.descriptor.class = newExpenseClass()
	return ExpenseClassCategorySelector{descriptor: this.descriptor}
}

type AssetClassCategorySelector struct{ descriptor AccountDescriptor }

func (this AssetClassCategorySelector) Cash() CashAssetCategoryGroupSelector {
	this.descriptor.category = newCashAssetCategory()
	return CashAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) Investment() InvestmentAssetCategoryGroupSelector {
	this.descriptor.category = newInvestmentAssetCategory()
	return InvestmentAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) Retirement() RetirementAssetCategoryGroupSelector {
	this.descriptor.category = newRetirementAssetCategory()
	return RetirementAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) Medical() MedicalAssetCategoryGroupSelector {
	this.descriptor.category = newMedicalAssetCategory()
	return MedicalAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) Education() EducationAssetCategoryGroupSelector {
	this.descriptor.category = newEducationAssetCategory()
	return EducationAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) RealEstate() RealEstateAssetCategoryGroupSelector {
	this.descriptor.category = newRealEstateAssetCategory()
	return RealEstateAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) Vehicle() VehicleAssetCategoryGroupSelector {
	this.descriptor.category = newVehicleAssetCategory()
	return VehicleAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) Household() HouseholdAssetCategoryGroupSelector {
	this.descriptor.category = newHouseholdAssetCategory()
	return HouseholdAssetCategoryGroupSelector{descriptor: this.descriptor}
}

func (this AssetClassCategorySelector) Personal() PersonalAssetCategoryGroupSelector {
	this.descriptor.category = newPersonalAssetCategory()
	return PersonalAssetCategoryGroupSelector{descriptor: this.descriptor}
}

type LiabilityClassCategorySelector struct{ descriptor AccountDescriptor }

func (this LiabilityClassCategorySelector) Bill() BillLiabilityCategoryGroupSelector {
	this.descriptor.category = newBillLiabilityCategory()
	return BillLiabilityCategoryGroupSelector{descriptor: this.descriptor}
}

func (this LiabilityClassCategorySelector) Credit() CreditLiabilityCategoryGroupSelector {
	this.descriptor.category = newCreditLiabilityCategory()
	return CreditLiabilityCategoryGroupSelector{descriptor: this.descriptor}
}

func (this LiabilityClassCategorySelector) Loan() LoanLiabilityCategoryGroupSelector {
	this.descriptor.category = newLoanLiabilityCategory()
	return LoanLiabilityCategoryGroupSelector{descriptor: this.descriptor}
}

func (this LiabilityClassCategorySelector) Tax() TaxLiabilityCategoryGroupSelector {
	this.descriptor.category = newTaxLiabilityCategory()
	return TaxLiabilityCategoryGroupSelector{descriptor: this.descriptor}
}

type EquityClassCategorySelector struct{ descriptor AccountDescriptor }

func (this EquityClassCategorySelector) Gift() GiftEquityCategoryGroupSelector {
	this.descriptor.category = newGiftEquityCategory()
	return GiftEquityCategoryGroupSelector{descriptor: this.descriptor}
}

func (this EquityClassCategorySelector) Capital() CapitalEquityCategoryGroupSelector {
	this.descriptor.category = newCapitalEquityCategory()
	return CapitalEquityCategoryGroupSelector{descriptor: this.descriptor}
}

func (this EquityClassCategorySelector) Unrealized() UnrealizedEquityCategoryGroupSelector {
	this.descriptor.category = newUnrealizedEquityCategory()
	return UnrealizedEquityCategoryGroupSelector{descriptor: this.descriptor}
}

type RevenueClassCategorySelector struct{ descriptor AccountDescriptor }

func (this RevenueClassCategorySelector) Employer() EmployerRevenueCategoryGroupSelector {
	this.descriptor.category = newEmployerRevenueCategory()
	return EmployerRevenueCategoryGroupSelector{descriptor: this.descriptor}
}

func (this RevenueClassCategorySelector) Earned() EarnedRevenueCategoryGroupSelector {
	this.descriptor.category = newEarnedRevenueCategory()
	return EarnedRevenueCategoryGroupSelector{descriptor: this.descriptor}
}

func (this RevenueClassCategorySelector) Passive() PassiveRevenueCategoryGroupSelector {
	this.descriptor.category = newPassiveRevenueCategory()
	return PassiveRevenueCategoryGroupSelector{descriptor: this.descriptor}
}

func (this RevenueClassCategorySelector) Government() GovernmentRevenueCategoryGroupSelector {
	this.descriptor.category = newGovernmentRevenueCategory()
	return GovernmentRevenueCategoryGroupSelector{descriptor: this.descriptor}
}

type ExpenseClassCategorySelector struct{ descriptor AccountDescriptor }

func (this ExpenseClassCategorySelector) Housing() HousingExpenseCategoryGroupSelector {
	this.descriptor.category = newHousingExpenseCategory()
	return HousingExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Household() HouseholdExpenseCategoryGroupSelector {
	this.descriptor.category = newHouseholdExpenseCategory()
	return HouseholdExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Food() FoodExpenseCategoryGroupSelector {
	this.descriptor.category = newFoodExpenseCategory()
	return FoodExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Transportation() TransportationExpenseCategoryGroupSelector {
	this.descriptor.category = newTransportationExpenseCategory()
	return TransportationExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Health() HealthExpenseCategoryGroupSelector {
	this.descriptor.category = newHealthExpenseCategory()
	return HealthExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Pet() PetExpenseCategoryGroupSelector {
	this.descriptor.category = newPetExpenseCategory()
	return PetExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Education() EducationExpenseCategoryGroupSelector {
	this.descriptor.category = newEducationExpenseCategory()
	return EducationExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Personal() PersonalExpenseCategoryGroupSelector {
	this.descriptor.category = newPersonalExpenseCategory()
	return PersonalExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Work() WorkExpenseCategoryGroupSelector {
	this.descriptor.category = newWorkExpenseCategory()
	return WorkExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Professional() ProfessionalExpenseCategoryGroupSelector {
	this.descriptor.category = newProfessionalExpenseCategory()
	return ProfessionalExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Government() GovernmentExpenseCategoryGroupSelector {
	this.descriptor.category = newGovernmentExpenseCategory()
	return GovernmentExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Financial() FinancialExpenseCategoryGroupSelector {
	this.descriptor.category = newFinancialExpenseCategory()
	return FinancialExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Gift() GiftExpenseCategoryGroupSelector {
	this.descriptor.category = newGiftExpenseCategory()
	return GiftExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Charitable() CharitableExpenseCategoryGroupSelector {
	this.descriptor.category = newCharitableExpenseCategory()
	return CharitableExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

func (this ExpenseClassCategorySelector) Miscellaneous() MiscellaneousExpenseCategoryGroupSelector {
	this.descriptor.category = newMiscellaneousExpenseCategory()
	return MiscellaneousExpenseCategoryGroupSelector{descriptor: this.descriptor}
}

type CashAssetCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this CashAssetCategoryGroupSelector) Petty() AccountFinalizer {
	this.descriptor.group = newPetGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CashAssetCategoryGroupSelector) Receivable() AccountFinalizer {
	this.descriptor.group = newReceivableGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CashAssetCategoryGroupSelector) Escrow() AccountFinalizer {
	this.descriptor.group = newEscrowGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CashAssetCategoryGroupSelector) Checking() AccountFinalizer {
	this.descriptor.group = newCheckingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CashAssetCategoryGroupSelector) Market() AccountFinalizer {
	this.descriptor.group = newMarketGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CashAssetCategoryGroupSelector) Saving() AccountFinalizer {
	this.descriptor.group = newSavingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CashAssetCategoryGroupSelector) CD() AccountFinalizer {
	this.descriptor.group = newCDGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CashAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type InvestmentAssetCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this InvestmentAssetCategoryGroupSelector) Stock() AccountFinalizer {
	this.descriptor.group = newStockGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this InvestmentAssetCategoryGroupSelector) Bond() AccountFinalizer {
	this.descriptor.group = newBondGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this InvestmentAssetCategoryGroupSelector) Pooled() AccountFinalizer {
	this.descriptor.group = newPooledGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this InvestmentAssetCategoryGroupSelector) Derivative() AccountFinalizer {
	this.descriptor.group = newDerivativeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this InvestmentAssetCategoryGroupSelector) Metal() AccountFinalizer {
	this.descriptor.group = newMetalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this InvestmentAssetCategoryGroupSelector) Crypto() AccountFinalizer {
	this.descriptor.group = newCryptoGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this InvestmentAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type RetirementAssetCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this RetirementAssetCategoryGroupSelector) Insurance() AccountFinalizer {
	this.descriptor.group = newInsuranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this RetirementAssetCategoryGroupSelector) Plan() AccountFinalizer {
	this.descriptor.group = newPlanGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this RetirementAssetCategoryGroupSelector) Arrangement() AccountFinalizer {
	this.descriptor.group = newArrangementGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this RetirementAssetCategoryGroupSelector) Foreign() AccountFinalizer {
	this.descriptor.group = newForeignGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this RetirementAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type MedicalAssetCategoryGroupSelector struct {
	descriptor AccountDescriptor
	CashAssetCategoryGroupSelector
	InvestmentAssetCategoryGroupSelector
}

func (this MedicalAssetCategoryGroupSelector) Plan() AccountFinalizer {
	this.descriptor.group = newPlanGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this MedicalAssetCategoryGroupSelector) Trust() AccountFinalizer {
	this.descriptor.group = newTrustGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this MedicalAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type EducationAssetCategoryGroupSelector struct {
	descriptor AccountDescriptor
	CashAssetCategoryGroupSelector
	InvestmentAssetCategoryGroupSelector
}

func (this EducationAssetCategoryGroupSelector) Plan() AccountFinalizer {
	this.descriptor.group = newPlanGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationAssetCategoryGroupSelector) Trust() AccountFinalizer {
	this.descriptor.group = newTrustGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type RealEstateAssetCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this RealEstateAssetCategoryGroupSelector) Residential() AccountFinalizer {
	this.descriptor.group = newResidentialGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this RealEstateAssetCategoryGroupSelector) Undeveloped() AccountFinalizer {
	this.descriptor.group = newUndevelopedGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this RealEstateAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type VehicleAssetCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this VehicleAssetCategoryGroupSelector) Road() AccountFinalizer {
	this.descriptor.group = newRoadGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this VehicleAssetCategoryGroupSelector) Recreational() AccountFinalizer {
	this.descriptor.group = newRecreationalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this VehicleAssetCategoryGroupSelector) Land() AccountFinalizer {
	this.descriptor.group = newLandGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this VehicleAssetCategoryGroupSelector) Water() AccountFinalizer {
	this.descriptor.group = newWaterGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this VehicleAssetCategoryGroupSelector) Air() AccountFinalizer {
	this.descriptor.group = newAirGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this VehicleAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type HouseholdAssetCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this HouseholdAssetCategoryGroupSelector) Furniture() AccountFinalizer {
	this.descriptor.group = newFurnitureGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdAssetCategoryGroupSelector) Electronic() AccountFinalizer {
	this.descriptor.group = newElectronicGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdAssetCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type PersonalAssetCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this PersonalAssetCategoryGroupSelector) Intellectual() AccountFinalizer {
	this.descriptor.group = newIntellectualGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalAssetCategoryGroupSelector) Jewelry() AccountFinalizer {
	this.descriptor.group = newJewelryGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalAssetCategoryGroupSelector) Collectible() AccountFinalizer {
	this.descriptor.group = newCollectibleGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalAssetCategoryGroupSelector) Tool() AccountFinalizer {
	this.descriptor.group = newToolGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalAssetCategoryGroupSelector) Electronic() AccountFinalizer {
	this.descriptor.group = newElectronicGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalAssetCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalAssetCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type BillLiabilityCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this BillLiabilityCategoryGroupSelector) Principal() AccountFinalizer {
	this.descriptor.group = newPrincipalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this BillLiabilityCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this BillLiabilityCategoryGroupSelector) Fee() AccountFinalizer {
	this.descriptor.group = newFeeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type CreditLiabilityCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this CreditLiabilityCategoryGroupSelector) Principal() AccountFinalizer {
	this.descriptor.group = newPrincipalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CreditLiabilityCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CreditLiabilityCategoryGroupSelector) Fee() AccountFinalizer {
	this.descriptor.group = newFeeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type LoanLiabilityCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this LoanLiabilityCategoryGroupSelector) Principal() AccountFinalizer {
	this.descriptor.group = newPrincipalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this LoanLiabilityCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this LoanLiabilityCategoryGroupSelector) Fee() AccountFinalizer {
	this.descriptor.group = newFeeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type TaxLiabilityCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this TaxLiabilityCategoryGroupSelector) Foreign() AccountFinalizer {
	this.descriptor.group = newForeignGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TaxLiabilityCategoryGroupSelector) Federal() AccountFinalizer {
	this.descriptor.group = newFederalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TaxLiabilityCategoryGroupSelector) State() AccountFinalizer {
	this.descriptor.group = newStateGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TaxLiabilityCategoryGroupSelector) Local() AccountFinalizer {
	this.descriptor.group = newLocalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type GiftEquityCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this GiftEquityCategoryGroupSelector) Gain() AccountFinalizer {
	this.descriptor.group = newGainGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GiftEquityCategoryGroupSelector) Loss() AccountFinalizer {
	this.descriptor.group = newLossGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type CapitalEquityCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this CapitalEquityCategoryGroupSelector) Gain() AccountFinalizer {
	this.descriptor.group = newGainGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CapitalEquityCategoryGroupSelector) Loss() AccountFinalizer {
	this.descriptor.group = newLossGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type UnrealizedEquityCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this UnrealizedEquityCategoryGroupSelector) Gain() AccountFinalizer {
	this.descriptor.group = newGainGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this UnrealizedEquityCategoryGroupSelector) Loss() AccountFinalizer {
	this.descriptor.group = newLossGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type EmployerRevenueCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this EmployerRevenueCategoryGroupSelector) Wage() AccountFinalizer {
	this.descriptor.group = newWageGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Overtime() AccountFinalizer {
	this.descriptor.group = newOvertimeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Tip() AccountFinalizer {
	this.descriptor.group = newTipGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Salary() AccountFinalizer {
	this.descriptor.group = newSalaryGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Stock() AccountFinalizer {
	this.descriptor.group = newStockGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Severance() AccountFinalizer {
	this.descriptor.group = newSeveranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Bonus() AccountFinalizer {
	this.descriptor.group = newBonusGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Reimbursement() AccountFinalizer {
	this.descriptor.group = newReimbursementGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Commission() AccountFinalizer {
	this.descriptor.group = newCommissionGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EmployerRevenueCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type EarnedRevenueCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this EarnedRevenueCategoryGroupSelector) Contract() AccountFinalizer {
	this.descriptor.group = newContractGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EarnedRevenueCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EarnedRevenueCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type PassiveRevenueCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this PassiveRevenueCategoryGroupSelector) Insurance() AccountFinalizer {
	this.descriptor.group = newInsuranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PassiveRevenueCategoryGroupSelector) Crypto() AccountFinalizer {
	this.descriptor.group = newCryptoGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PassiveRevenueCategoryGroupSelector) Dividend() AccountFinalizer {
	this.descriptor.group = newDividendGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PassiveRevenueCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PassiveRevenueCategoryGroupSelector) Royalty() AccountFinalizer {
	this.descriptor.group = newRoyaltyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PassiveRevenueCategoryGroupSelector) Trust() AccountFinalizer {
	this.descriptor.group = newTrustGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PassiveRevenueCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PassiveRevenueCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type GovernmentRevenueCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this GovernmentRevenueCategoryGroupSelector) Tax() AccountFinalizer {
	this.descriptor.group = newTaxGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentRevenueCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentRevenueCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type HousingExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this HousingExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Insurance() AccountFinalizer {
	this.descriptor.group = newInsuranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Rent() AccountFinalizer {
	this.descriptor.group = newRentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Maintenance() AccountFinalizer {
	this.descriptor.group = newMaintenanceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Renovation() AccountFinalizer {
	this.descriptor.group = newRenovationGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Electric() AccountFinalizer {
	this.descriptor.group = newElectricGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Gas() AccountFinalizer {
	this.descriptor.group = newGasGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Internet() AccountFinalizer {
	this.descriptor.group = newInternetGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Phone() AccountFinalizer {
	this.descriptor.group = newPhoneGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Trash() AccountFinalizer {
	this.descriptor.group = newTrashGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Water() AccountFinalizer {
	this.descriptor.group = newWaterGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Sewer() AccountFinalizer {
	this.descriptor.group = newSewerGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HousingExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type HouseholdExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this HouseholdExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdExpenseCategoryGroupSelector) Insurance() AccountFinalizer {
	this.descriptor.group = newInsuranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdExpenseCategoryGroupSelector) Furniture() AccountFinalizer {
	this.descriptor.group = newFurnitureGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdExpenseCategoryGroupSelector) Electronic() AccountFinalizer {
	this.descriptor.group = newElectronicGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdExpenseCategoryGroupSelector) Supply() AccountFinalizer {
	this.descriptor.group = newSupplyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdExpenseCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HouseholdExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type FoodExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this FoodExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Fruit() AccountFinalizer {
	this.descriptor.group = newFruitGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Vegetable() AccountFinalizer {
	this.descriptor.group = newVegetableGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Grain() AccountFinalizer {
	this.descriptor.group = newGrainGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Protein() AccountFinalizer {
	this.descriptor.group = newProteinGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Dairy() AccountFinalizer {
	this.descriptor.group = newDairyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Oil() AccountFinalizer {
	this.descriptor.group = newOilGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Supplement() AccountFinalizer {
	this.descriptor.group = newSupplementGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Beverage() AccountFinalizer {
	this.descriptor.group = newBeverageGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Baked() AccountFinalizer {
	this.descriptor.group = newBakedGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Frozen() AccountFinalizer {
	this.descriptor.group = newFrozenGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Canned() AccountFinalizer {
	this.descriptor.group = newCannedGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Boxed() AccountFinalizer {
	this.descriptor.group = newBoxedGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Jarred() AccountFinalizer {
	this.descriptor.group = newJarredGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Pet() AccountFinalizer {
	this.descriptor.group = newPetGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Baby() AccountFinalizer {
	this.descriptor.group = newBabyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Consumable() AccountFinalizer {
	this.descriptor.group = newConsumableGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) FastFood() AccountFinalizer {
	this.descriptor.group = newFastFoodGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Delivery() AccountFinalizer {
	this.descriptor.group = newDeliveryGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) DiningOut() AccountFinalizer {
	this.descriptor.group = newDiningOutGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FoodExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type TransportationExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this TransportationExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Insurance() AccountFinalizer {
	this.descriptor.group = newInsuranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Fuel() AccountFinalizer {
	this.descriptor.group = newFuelGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Maintenance() AccountFinalizer {
	this.descriptor.group = newMaintenanceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Repair() AccountFinalizer {
	this.descriptor.group = newRepairGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Shared() AccountFinalizer {
	this.descriptor.group = newSharedGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Fee() AccountFinalizer {
	this.descriptor.group = newFeeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this TransportationExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type HealthExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this HealthExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Insurance() AccountFinalizer {
	this.descriptor.group = newInsuranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Medical() AccountFinalizer {
	this.descriptor.group = newMedicalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Dental() AccountFinalizer {
	this.descriptor.group = newDentalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Vision() AccountFinalizer {
	this.descriptor.group = newVisionGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Hygiene() AccountFinalizer {
	this.descriptor.group = newHygieneGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Supply() AccountFinalizer {
	this.descriptor.group = newSupplyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this HealthExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type PetExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this PetExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Insurance() AccountFinalizer {
	this.descriptor.group = newInsuranceGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Vet() AccountFinalizer {
	this.descriptor.group = newVetGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Hygiene() AccountFinalizer {
	this.descriptor.group = newHygieneGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Grooming() AccountFinalizer {
	this.descriptor.group = newGroomingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Supply() AccountFinalizer {
	this.descriptor.group = newSupplyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PetExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type EducationExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this EducationExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Tuition() AccountFinalizer {
	this.descriptor.group = newTuitionGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Credential() AccountFinalizer {
	this.descriptor.group = newCredentialGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Reading() AccountFinalizer {
	this.descriptor.group = newReadingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Travel() AccountFinalizer {
	this.descriptor.group = newTravelGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Electronic() AccountFinalizer {
	this.descriptor.group = newElectronicGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Supply() AccountFinalizer {
	this.descriptor.group = newSupplyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this EducationExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type PersonalExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this PersonalExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Grooming() AccountFinalizer {
	this.descriptor.group = newGroomingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Entertainment() AccountFinalizer {
	this.descriptor.group = newEntertainmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Clothing() AccountFinalizer {
	this.descriptor.group = newClothingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Reading() AccountFinalizer {
	this.descriptor.group = newReadingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Travel() AccountFinalizer {
	this.descriptor.group = newTravelGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Electronic() AccountFinalizer {
	this.descriptor.group = newElectronicGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Supply() AccountFinalizer {
	this.descriptor.group = newSupplyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this PersonalExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type WorkExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this WorkExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Tuition() AccountFinalizer {
	this.descriptor.group = newTuitionGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Credential() AccountFinalizer {
	this.descriptor.group = newCredentialGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Clothing() AccountFinalizer {
	this.descriptor.group = newClothingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Reading() AccountFinalizer {
	this.descriptor.group = newReadingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Travel() AccountFinalizer {
	this.descriptor.group = newTravelGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Electronic() AccountFinalizer {
	this.descriptor.group = newElectronicGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Supply() AccountFinalizer {
	this.descriptor.group = newSupplyGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Equipment() AccountFinalizer {
	this.descriptor.group = newEquipmentGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this WorkExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type ProfessionalExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this ProfessionalExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Legal() AccountFinalizer {
	this.descriptor.group = newLegalGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Accounting() AccountFinalizer {
	this.descriptor.group = newAccountingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Creative() AccountFinalizer {
	this.descriptor.group = newCreativeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Marketing() AccountFinalizer {
	this.descriptor.group = newMarketingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Consulting() AccountFinalizer {
	this.descriptor.group = newConsultingGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Administrative() AccountFinalizer {
	this.descriptor.group = newAdministrativeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this ProfessionalExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type GovernmentExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this GovernmentExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentExpenseCategoryGroupSelector) Tax() AccountFinalizer {
	this.descriptor.group = newTaxGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentExpenseCategoryGroupSelector) Fee() AccountFinalizer {
	this.descriptor.group = newFeeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentExpenseCategoryGroupSelector) Fine() AccountFinalizer {
	this.descriptor.group = newFineGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentExpenseCategoryGroupSelector) Charge() AccountFinalizer {
	this.descriptor.group = newChargeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GovernmentExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type FinancialExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this FinancialExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FinancialExpenseCategoryGroupSelector) Fee() AccountFinalizer {
	this.descriptor.group = newFeeGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FinancialExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this FinancialExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type GiftExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this GiftExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GiftExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this GiftExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type CharitableExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this CharitableExpenseCategoryGroupSelector) Interest() AccountFinalizer {
	this.descriptor.group = newInterestGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CharitableExpenseCategoryGroupSelector) Recurring() AccountFinalizer {
	this.descriptor.group = newRecurringGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this CharitableExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type MiscellaneousExpenseCategoryGroupSelector struct{ descriptor AccountDescriptor }

func (this MiscellaneousExpenseCategoryGroupSelector) Unaccounted() AccountFinalizer {
	this.descriptor.group = newUnaccountedGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

func (this MiscellaneousExpenseCategoryGroupSelector) Other() AccountFinalizer {
	this.descriptor.group = newOtherGroup()
	return AccountFinalizer{descriptor: this.descriptor}
}

type AccountFinalizer struct{ descriptor AccountDescriptor }

func (this AccountFinalizer) Describe(container, item, id string) AccountDescriptor {
	this.descriptor.container = container
	this.descriptor.item = item
	this.descriptor.id = id

	return this.descriptor
}

type AccountDescriptor struct {
	class    class
	category category
	group    group

	container string
	item      string
	id        string
}
