import { OrderStatus } from '../entities/order.entity';

export class UpdateOrderDto {
  amount: number;
  credit_card_number: string;
  credit_card_name: string;
  status: OrderStatus;
}
