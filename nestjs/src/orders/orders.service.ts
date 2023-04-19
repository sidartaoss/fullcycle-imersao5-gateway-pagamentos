import { Inject, Injectable } from '@nestjs/common';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { InjectModel } from '@nestjs/sequelize';
import { EmptyResultError } from 'sequelize';
import { AccountStorageService } from '../accounts/account-storage/account-storage.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  constructor(
    @InjectModel(Order)
    private orderModel: typeof Order,
    private accountStorageService: AccountStorageService,
    @Inject('KAFKA_PRODUCER')
    private kafkaProducer: Producer,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    const order = await this.orderModel.create({
      ...createOrderDto,
      account_id: this.accountStorageService.account.id,
    });
    await this.kafkaProducer.send({
      topic: 'transactions',
      messages: [
        {
          key: 'transactions',
          value: JSON.stringify({
            id: order.id,
            status: order.status,
            credit_card_number: order.credit_card_number,
            credit_card_name: order.credit_card_name,
            credit_card_expiration_month:
              createOrderDto.credit_card_expiration_month,
            credit_card_expiration_year:
              createOrderDto.credit_card_expiration_year,
            credit_card_cvv: createOrderDto.credit_card_cvv,
            amount: order.amount,
            account_id: order.account_id,
            created_at: order.createdAt,
            updated_at: order.updatedAt,
          }),
        },
      ],
    });
    return order;
  }

  findAll() {
    return this.orderModel.findAll({
      where: {
        account_id: this.accountStorageService.account.id,
      },
    });
  }

  findOneUsingAccount(id: string) {
    return this.orderModel.findOne({
      where: {
        id,
        account_id: this.accountStorageService.account.id,
      },
      rejectOnEmpty: new EmptyResultError(`Order with ID ${id} not found`),
    });
  }

  findOne(id: string) {
    return this.orderModel.findByPk(id);
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const account = this.accountStorageService.account;
    const order = await (account
      ? this.findOneUsingAccount(id)
      : this.findOne(id));
    return await order.update(updateOrderDto);
  }

  async remove(id: string) {
    const order = await this.findOneUsingAccount(id);
    await order.destroy();
  }
}
