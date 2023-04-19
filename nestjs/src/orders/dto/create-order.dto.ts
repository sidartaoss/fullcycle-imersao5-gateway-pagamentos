export class CreateOrderDto {
  credit_card_number: string;
  credit_card_name: string;
  credit_card_expiration_month: number;
  credit_card_expiration_year: number;
  credit_card_cvv: number;
  amount: number;
}
