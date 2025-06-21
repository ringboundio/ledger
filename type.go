package ledger

import "fmt"

type class interface {
	classID() string
	name() string
}

type baseClass _type

func newBaseClass(name string) baseClass {
	return baseClass{
		id:    cat("class", name),
		short: name,
	}
}

func (this baseClass) classID() string { return this.id }
func (this baseClass) name() string    { return this.short }

func newAssetClass() baseClass     { return newAssetSymbol().toBaseClass() }
func newLiabilityClass() baseClass { return newLiabilitySymbol().toBaseClass() }
func newEquityClass() baseClass    { return newEquitySymbol().toBaseClass() }
func newRevenueClass() baseClass   { return newRevenueSymbol().toBaseClass() }
func newExpenseClass() baseClass   { return newExpenseSymbol().toBaseClass() }

type category interface {
	categoryID() string
	name() string
}

type baseCategory _type

func newBaseCategory(name string) baseCategory {
	return baseCategory{
		id:    cat("category", name),
		short: name,
	}
}

func (this baseCategory) categoryID() string { return this.id }
func (this baseCategory) name() string       { return this.short }

func newCashAssetCategory() baseCategory       { return newCashSymbol().toBaseCategory() }
func newInvestmentAssetCategory() baseCategory { return newInvestmentSymbol().toBaseCategory() }
func newRetirementAssetCategory() baseCategory { return newRetirementSymbol().toBaseCategory() }
func newMedicalAssetCategory() baseCategory    { return newMedicalSymbol().toBaseCategory() }
func newEducationAssetCategory() baseCategory  { return newEducationSymbol().toBaseCategory() }
func newRealEstateAssetCategory() baseCategory { return newRealEstateSymbol().toBaseCategory() }
func newVehicleAssetCategory() baseCategory    { return newVehicleSymbol().toBaseCategory() }
func newHouseholdAssetCategory() baseCategory  { return newHouseholdSymbol().toBaseCategory() }
func newPersonalAssetCategory() baseCategory   { return newPersonalSymbol().toBaseCategory() }

func newBillLiabilityCategory() baseCategory   { return newBillSymbol().toBaseCategory() }
func newCreditLiabilityCategory() baseCategory { return newCreditSymbol().toBaseCategory() }
func newLoanLiabilityCategory() baseCategory   { return newLoanSymbol().toBaseCategory() }
func newTaxLiabilityCategory() baseCategory    { return newTaxSymbol().toBaseCategory() }

func newGiftEquityCategory() baseCategory       { return newGiftSymbol().toBaseCategory() }
func newCapitalEquityCategory() baseCategory    { return newCapitalSymbol().toBaseCategory() }
func newUnrealizedEquityCategory() baseCategory { return newUnrealizedSymbol().toBaseCategory() }

func newEmployerRevenueCategory() baseCategory   { return newEmployerSymbol().toBaseCategory() }
func newEarnedRevenueCategory() baseCategory     { return newEarnedSymbol().toBaseCategory() }
func newPassiveRevenueCategory() baseCategory    { return newPassiveSymbol().toBaseCategory() }
func newGovernmentRevenueCategory() baseCategory { return newGovernmentSymbol().toBaseCategory() }

func newHousingExpenseCategory() baseCategory   { return newHousingSymbol().toBaseCategory() }
func newHouseholdExpenseCategory() baseCategory { return newHouseholdSymbol().toBaseCategory() }
func newFoodExpenseCategory() baseCategory      { return newFoodSymbol().toBaseCategory() }
func newTransportationExpenseCategory() baseCategory {
	return newTransportationSymbol().toBaseCategory()
}
func newHealthExpenseCategory() baseCategory        { return newHealthSymbol().toBaseCategory() }
func newPetExpenseCategory() baseCategory           { return newPetSymbol().toBaseCategory() }
func newEducationExpenseCategory() baseCategory     { return newEducationSymbol().toBaseCategory() }
func newPersonalExpenseCategory() baseCategory      { return newPersonalSymbol().toBaseCategory() }
func newWorkExpenseCategory() baseCategory          { return newWorkSymbol().toBaseCategory() }
func newProfessionalExpenseCategory() baseCategory  { return newProfessionalSymbol().toBaseCategory() }
func newGovernmentExpenseCategory() baseCategory    { return newGovernmentSymbol().toBaseCategory() }
func newFinancialExpenseCategory() baseCategory     { return newFinancialSymbol().toBaseCategory() }
func newGiftExpenseCategory() baseCategory          { return newGiftSymbol().toBaseCategory() }
func newCharitableExpenseCategory() baseCategory    { return newCharitableSymbol().toBaseCategory() }
func newMiscellaneousExpenseCategory() baseCategory { return newMiscellaneousSymbol().toBaseCategory() }

type group interface {
	groupID() string
	name() string
}

type baseGroup _type

func newBaseGroup(name string) baseGroup {
	return baseGroup{
		id:    cat("group", name),
		short: name,
	}
}

func (this baseGroup) groupID() string { return this.id }
func (this baseGroup) name() string    { return this.short }

func newOtherGroup() baseGroup          { return newOtherSymbol().toBaseGroup() }
func newPettyGroup() baseGroup          { return newPettySymbol().toBaseGroup() }
func newReceivableGroup() baseGroup     { return newReceivableSymbol().toBaseGroup() }
func newEscrowGroup() baseGroup         { return newEscrowSymbol().toBaseGroup() }
func newCheckingGroup() baseGroup       { return newCheckingSymbol().toBaseGroup() }
func newMarketGroup() baseGroup         { return newMarketSymbol().toBaseGroup() }
func newSavingGroup() baseGroup         { return newSavingSymbol().toBaseGroup() }
func newCDGroup() baseGroup             { return newCDSymbol().toBaseGroup() }
func newStockGroup() baseGroup          { return newStockSymbol().toBaseGroup() }
func newBondGroup() baseGroup           { return newBondSymbol().toBaseGroup() }
func newPooledGroup() baseGroup         { return newPooledSymbol().toBaseGroup() }
func newDerivativeGroup() baseGroup     { return newDerivativeSymbol().toBaseGroup() }
func newMetalGroup() baseGroup          { return newMetalSymbol().toBaseGroup() }
func newCryptoGroup() baseGroup         { return newCryptoSymbol().toBaseGroup() }
func newInsuranceGroup() baseGroup      { return newInsuranceSymbol().toBaseGroup() }
func newPlanGroup() baseGroup           { return newPlanSymbol().toBaseGroup() }
func newArrangementGroup() baseGroup    { return newArrangementSymbol().toBaseGroup() }
func newForeignGroup() baseGroup        { return newForeignSymbol().toBaseGroup() }
func newTrustGroup() baseGroup          { return newTrustSymbol().toBaseGroup() }
func newResidentialGroup() baseGroup    { return newResidentialSymbol().toBaseGroup() }
func newUndevelopedGroup() baseGroup    { return newUndevelopedSymbol().toBaseGroup() }
func newRoadGroup() baseGroup           { return newRoadSymbol().toBaseGroup() }
func newRecreationalGroup() baseGroup   { return newRecreationalSymbol().toBaseGroup() }
func newLandGroup() baseGroup           { return newLandSymbol().toBaseGroup() }
func newWaterGroup() baseGroup          { return newWaterSymbol().toBaseGroup() }
func newAirGroup() baseGroup            { return newAirSymbol().toBaseGroup() }
func newFurnitureGroup() baseGroup      { return newFurnitureSymbol().toBaseGroup() }
func newElectronicGroup() baseGroup     { return newElectronicSymbol().toBaseGroup() }
func newEquipmentGroup() baseGroup      { return newEquipmentSymbol().toBaseGroup() }
func newIntellectualGroup() baseGroup   { return newIntellectualSymbol().toBaseGroup() }
func newJewelryGroup() baseGroup        { return newJewelrySymbol().toBaseGroup() }
func newCollectibleGroup() baseGroup    { return newCollectibleSymbol().toBaseGroup() }
func newToolGroup() baseGroup           { return newToolSymbol().toBaseGroup() }
func newPrincipalGroup() baseGroup      { return newPrincipalSymbol().toBaseGroup() }
func newInterestGroup() baseGroup       { return newInterestSymbol().toBaseGroup() }
func newFeeGroup() baseGroup            { return newFeeSymbol().toBaseGroup() }
func newFederalGroup() baseGroup        { return newFederalSymbol().toBaseGroup() }
func newStateGroup() baseGroup          { return newStateSymbol().toBaseGroup() }
func newLocalGroup() baseGroup          { return newLocalSymbol().toBaseGroup() }
func newGainGroup() baseGroup           { return newGainSymbol().toBaseGroup() }
func newLossGroup() baseGroup           { return newLossSymbol().toBaseGroup() }
func newWageGroup() baseGroup           { return newWageSymbol().toBaseGroup() }
func newOvertimeGroup() baseGroup       { return newOvertimeSymbol().toBaseGroup() }
func newTipGroup() baseGroup            { return newTipSymbol().toBaseGroup() }
func newSalaryGroup() baseGroup         { return newSalarySymbol().toBaseGroup() }
func newSeveranceGroup() baseGroup      { return newSeveranceSymbol().toBaseGroup() }
func newBonusGroup() baseGroup          { return newBonusSymbol().toBaseGroup() }
func newReimbursementGroup() baseGroup  { return newReimbursementSymbol().toBaseGroup() }
func newCommissionGroup() baseGroup     { return newCommissionSymbol().toBaseGroup() }
func newRecurringGroup() baseGroup      { return newRecurringSymbol().toBaseGroup() }
func newContractGroup() baseGroup       { return newContractSymbol().toBaseGroup() }
func newDividendGroup() baseGroup       { return newDividendSymbol().toBaseGroup() }
func newRoyaltyGroup() baseGroup        { return newRoyaltySymbol().toBaseGroup() }
func newTaxGroup() baseGroup            { return newTaxSymbol().toBaseGroup() }
func newRentGroup() baseGroup           { return newRentSymbol().toBaseGroup() }
func newMaintenanceGroup() baseGroup    { return newMaintenanceSymbol().toBaseGroup() }
func newRenovationGroup() baseGroup     { return newRenovationSymbol().toBaseGroup() }
func newElectricGroup() baseGroup       { return newElectricSymbol().toBaseGroup() }
func newGasGroup() baseGroup            { return newGasSymbol().toBaseGroup() }
func newInternetGroup() baseGroup       { return newInternetSymbol().toBaseGroup() }
func newPhoneGroup() baseGroup          { return newPhoneSymbol().toBaseGroup() }
func newTrashGroup() baseGroup          { return newTrashSymbol().toBaseGroup() }
func newSewerGroup() baseGroup          { return newSewerSymbol().toBaseGroup() }
func newSupplyGroup() baseGroup         { return newSupplySymbol().toBaseGroup() }
func newFruitGroup() baseGroup          { return newFruitSymbol().toBaseGroup() }
func newVegetableGroup() baseGroup      { return newVegetableSymbol().toBaseGroup() }
func newGrainGroup() baseGroup          { return newGrainSymbol().toBaseGroup() }
func newProteinGroup() baseGroup        { return newProteinSymbol().toBaseGroup() }
func newDairyGroup() baseGroup          { return newDairySymbol().toBaseGroup() }
func newOilGroup() baseGroup            { return newOilSymbol().toBaseGroup() }
func newSupplementGroup() baseGroup     { return newSupplementSymbol().toBaseGroup() }
func newBeverageGroup() baseGroup       { return newBeverageSymbol().toBaseGroup() }
func newBakedGroup() baseGroup          { return newBakedSymbol().toBaseGroup() }
func newFrozenGroup() baseGroup         { return newFrozenSymbol().toBaseGroup() }
func newCannedGroup() baseGroup         { return newCannedSymbol().toBaseGroup() }
func newBoxedGroup() baseGroup          { return newBoxedSymbol().toBaseGroup() }
func newJarredGroup() baseGroup         { return newJarredSymbol().toBaseGroup() }
func newPetGroup() baseGroup            { return newPetSymbol().toBaseGroup() }
func newBabyGroup() baseGroup           { return newBabySymbol().toBaseGroup() }
func newConsumableGroup() baseGroup     { return newConsumableSymbol().toBaseGroup() }
func newFastFoodGroup() baseGroup       { return newFastFoodSymbol().toBaseGroup() }
func newDeliveryGroup() baseGroup       { return newDeliverySymbol().toBaseGroup() }
func newDiningOutGroup() baseGroup      { return newDiningOutSymbol().toBaseGroup() }
func newFuelGroup() baseGroup           { return newFuelSymbol().toBaseGroup() }
func newRepairGroup() baseGroup         { return newRepairSymbol().toBaseGroup() }
func newSharedGroup() baseGroup         { return newSharedSymbol().toBaseGroup() }
func newMedicalGroup() baseGroup        { return newMedicalSymbol().toBaseGroup() }
func newDentalGroup() baseGroup         { return newDentalSymbol().toBaseGroup() }
func newVisionGroup() baseGroup         { return newVisionSymbol().toBaseGroup() }
func newHygieneGroup() baseGroup        { return newHygieneSymbol().toBaseGroup() }
func newVetGroup() baseGroup            { return newVetSymbol().toBaseGroup() }
func newGroomingGroup() baseGroup       { return newGroomingSymbol().toBaseGroup() }
func newTuitionGroup() baseGroup        { return newTuitionSymbol().toBaseGroup() }
func newCredentialGroup() baseGroup     { return newCredentialSymbol().toBaseGroup() }
func newReadingGroup() baseGroup        { return newReadingSymbol().toBaseGroup() }
func newTravelGroup() baseGroup         { return newTravelSymbol().toBaseGroup() }
func newEntertainmentGroup() baseGroup  { return newEntertainmentSymbol().toBaseGroup() }
func newClothingGroup() baseGroup       { return newClothingSymbol().toBaseGroup() }
func newLegalGroup() baseGroup          { return newLegalSymbol().toBaseGroup() }
func newAccountingGroup() baseGroup     { return newAccountingSymbol().toBaseGroup() }
func newCreativeGroup() baseGroup       { return newCreativeSymbol().toBaseGroup() }
func newMarketingGroup() baseGroup      { return newMarketingSymbol().toBaseGroup() }
func newConsultingGroup() baseGroup     { return newConsultingSymbol().toBaseGroup() }
func newAdministrativeGroup() baseGroup { return newAdministrativeSymbol().toBaseGroup() }
func newFineGroup() baseGroup           { return newFineSymbol().toBaseGroup() }
func newChargeGroup() baseGroup         { return newChargeSymbol().toBaseGroup() }
func newUnaccountedGroup() baseGroup    { return newUnaccountedSymbol().toBaseGroup() }

type symbol struct{ name string }

func newSymbol(name string) symbol {
	return symbol{name: name}
}

func (this symbol) toBaseClass() baseClass       { return newBaseClass(this.name) }
func (this symbol) toBaseCategory() baseCategory { return newBaseCategory(this.name) }
func (this symbol) toBaseGroup() baseGroup       { return newBaseGroup(this.name) }

func newAssetSymbol() symbol          { return newSymbol("asset") }
func newLiabilitySymbol() symbol      { return newSymbol("liability") }
func newEquitySymbol() symbol         { return newSymbol("equity") }
func newRevenueSymbol() symbol        { return newSymbol("revenue") }
func newExpenseSymbol() symbol        { return newSymbol("expense") }
func newCashSymbol() symbol           { return newSymbol("cash") }
func newInvestmentSymbol() symbol     { return newSymbol("investment") }
func newRetirementSymbol() symbol     { return newSymbol("retirement") }
func newMedicalSymbol() symbol        { return newSymbol("medical") }
func newEducationSymbol() symbol      { return newSymbol("education") }
func newRealEstateSymbol() symbol     { return newSymbol("real estate") }
func newVehicleSymbol() symbol        { return newSymbol("vehicle") }
func newHouseholdSymbol() symbol      { return newSymbol("household") }
func newPersonalSymbol() symbol       { return newSymbol("personal") }
func newBillSymbol() symbol           { return newSymbol("bill") }
func newCreditSymbol() symbol         { return newSymbol("credit") }
func newLoanSymbol() symbol           { return newSymbol("loan") }
func newTaxSymbol() symbol            { return newSymbol("tax") }
func newGiftSymbol() symbol           { return newSymbol("gift") }
func newCapitalSymbol() symbol        { return newSymbol("capital") }
func newUnrealizedSymbol() symbol     { return newSymbol("unrealized") }
func newEmployerSymbol() symbol       { return newSymbol("employer") }
func newEarnedSymbol() symbol         { return newSymbol("earned") }
func newPassiveSymbol() symbol        { return newSymbol("passive") }
func newGovernmentSymbol() symbol     { return newSymbol("government") }
func newHousingSymbol() symbol        { return newSymbol("housing") }
func newFoodSymbol() symbol           { return newSymbol("food") }
func newTransportationSymbol() symbol { return newSymbol("transportation") }
func newHealthSymbol() symbol         { return newSymbol("health") }
func newPetSymbol() symbol            { return newSymbol("pet") }
func newWorkSymbol() symbol           { return newSymbol("work") }
func newProfessionalSymbol() symbol   { return newSymbol("professional") }
func newFinancialSymbol() symbol      { return newSymbol("financial") }
func newCharitableSymbol() symbol     { return newSymbol("charitable") }
func newMiscellaneousSymbol() symbol  { return newSymbol("miscellaneous") }
func newOtherSymbol() symbol          { return newSymbol("other") }
func newPettySymbol() symbol          { return newSymbol("petty") }
func newReceivableSymbol() symbol     { return newSymbol("receivable") }
func newEscrowSymbol() symbol         { return newSymbol("escrow") }
func newCheckingSymbol() symbol       { return newSymbol("checking") }
func newMarketSymbol() symbol         { return newSymbol("market") }
func newSavingSymbol() symbol         { return newSymbol("saving") }
func newCDSymbol() symbol             { return newSymbol("cd") }
func newStockSymbol() symbol          { return newSymbol("stock") }
func newBondSymbol() symbol           { return newSymbol("bond") }
func newPooledSymbol() symbol         { return newSymbol("pooled") }
func newDerivativeSymbol() symbol     { return newSymbol("derivative") }
func newMetalSymbol() symbol          { return newSymbol("metal") }
func newCryptoSymbol() symbol         { return newSymbol("crypto") }
func newInsuranceSymbol() symbol      { return newSymbol("insurance") }
func newPlanSymbol() symbol           { return newSymbol("plan") }
func newArrangementSymbol() symbol    { return newSymbol("arrangement") }
func newForeignSymbol() symbol        { return newSymbol("foreign") }
func newTrustSymbol() symbol          { return newSymbol("trust") }
func newResidentialSymbol() symbol    { return newSymbol("residential") }
func newUndevelopedSymbol() symbol    { return newSymbol("undeveloped") }
func newRoadSymbol() symbol           { return newSymbol("road") }
func newRecreationalSymbol() symbol   { return newSymbol("recreational") }
func newLandSymbol() symbol           { return newSymbol("land") }
func newWaterSymbol() symbol          { return newSymbol("water") }
func newAirSymbol() symbol            { return newSymbol("air") }
func newFurnitureSymbol() symbol      { return newSymbol("furniture") }
func newElectronicSymbol() symbol     { return newSymbol("electronic") }
func newEquipmentSymbol() symbol      { return newSymbol("equipment") }
func newIntellectualSymbol() symbol   { return newSymbol("intellectual") }
func newJewelrySymbol() symbol        { return newSymbol("jewelry") }
func newCollectibleSymbol() symbol    { return newSymbol("collectible") }
func newToolSymbol() symbol           { return newSymbol("tool") }
func newPrincipalSymbol() symbol      { return newSymbol("principal") }
func newInterestSymbol() symbol       { return newSymbol("interest") }
func newFeeSymbol() symbol            { return newSymbol("fee") }
func newFederalSymbol() symbol        { return newSymbol("federal") }
func newStateSymbol() symbol          { return newSymbol("state") }
func newLocalSymbol() symbol          { return newSymbol("local") }
func newGainSymbol() symbol           { return newSymbol("gain") }
func newLossSymbol() symbol           { return newSymbol("loss") }
func newWageSymbol() symbol           { return newSymbol("wage") }
func newOvertimeSymbol() symbol       { return newSymbol("overtime") }
func newTipSymbol() symbol            { return newSymbol("tip") }
func newSalarySymbol() symbol         { return newSymbol("salary") }
func newSeveranceSymbol() symbol      { return newSymbol("severance") }
func newBonusSymbol() symbol          { return newSymbol("bonus") }
func newReimbursementSymbol() symbol  { return newSymbol("reimbursement") }
func newCommissionSymbol() symbol     { return newSymbol("commission") }
func newRecurringSymbol() symbol      { return newSymbol("recurring") }
func newContractSymbol() symbol       { return newSymbol("contract") }
func newDividendSymbol() symbol       { return newSymbol("dividend") }
func newRoyaltySymbol() symbol        { return newSymbol("royalty") }
func newRentSymbol() symbol           { return newSymbol("rent") }
func newMaintenanceSymbol() symbol    { return newSymbol("maintenance") }
func newRenovationSymbol() symbol     { return newSymbol("renovation") }
func newElectricSymbol() symbol       { return newSymbol("electric") }
func newGasSymbol() symbol            { return newSymbol("gas") }
func newInternetSymbol() symbol       { return newSymbol("internet") }
func newPhoneSymbol() symbol          { return newSymbol("phone") }
func newTrashSymbol() symbol          { return newSymbol("trash") }
func newSewerSymbol() symbol          { return newSymbol("sewer") }
func newSupplySymbol() symbol         { return newSymbol("supply") }
func newFruitSymbol() symbol          { return newSymbol("fruit") }
func newVegetableSymbol() symbol      { return newSymbol("vegetable") }
func newGrainSymbol() symbol          { return newSymbol("grain") }
func newProteinSymbol() symbol        { return newSymbol("protein") }
func newDairySymbol() symbol          { return newSymbol("dairy") }
func newOilSymbol() symbol            { return newSymbol("oil") }
func newSupplementSymbol() symbol     { return newSymbol("supplement") }
func newBeverageSymbol() symbol       { return newSymbol("beverage") }
func newBakedSymbol() symbol          { return newSymbol("baked") }
func newFrozenSymbol() symbol         { return newSymbol("frozen") }
func newCannedSymbol() symbol         { return newSymbol("canned") }
func newBoxedSymbol() symbol          { return newSymbol("boxed") }
func newJarredSymbol() symbol         { return newSymbol("jarred") }
func newBabySymbol() symbol           { return newSymbol("baby") }
func newConsumableSymbol() symbol     { return newSymbol("consumable") }
func newFastFoodSymbol() symbol       { return newSymbol("fast food") }
func newDeliverySymbol() symbol       { return newSymbol("delivery") }
func newDiningOutSymbol() symbol      { return newSymbol("dining out") }
func newFuelSymbol() symbol           { return newSymbol("fuel") }
func newRepairSymbol() symbol         { return newSymbol("repair") }
func newSharedSymbol() symbol         { return newSymbol("shared") }
func newDentalSymbol() symbol         { return newSymbol("dental") }
func newVisionSymbol() symbol         { return newSymbol("vision") }
func newHygieneSymbol() symbol        { return newSymbol("hygiene") }
func newVetSymbol() symbol            { return newSymbol("vet") }
func newGroomingSymbol() symbol       { return newSymbol("grooming") }
func newTuitionSymbol() symbol        { return newSymbol("tuition") }
func newCredentialSymbol() symbol     { return newSymbol("credential") }
func newReadingSymbol() symbol        { return newSymbol("reading") }
func newTravelSymbol() symbol         { return newSymbol("travel") }
func newEntertainmentSymbol() symbol  { return newSymbol("entertainment") }
func newClothingSymbol() symbol       { return newSymbol("clothing") }
func newLegalSymbol() symbol          { return newSymbol("legal") }
func newAccountingSymbol() symbol     { return newSymbol("accounting") }
func newCreativeSymbol() symbol       { return newSymbol("creative") }
func newMarketingSymbol() symbol      { return newSymbol("marketing") }
func newConsultingSymbol() symbol     { return newSymbol("consulting") }
func newAdministrativeSymbol() symbol { return newSymbol("administrative") }
func newFineSymbol() symbol           { return newSymbol("fine") }
func newChargeSymbol() symbol         { return newSymbol("charge") }
func newUnaccountedSymbol() symbol    { return newSymbol("unaccounted") }

type _type struct {
	id, short string
}

func cat(base, name string) string {
	return fmt.Sprintf("%s:%s", base, name)
}
