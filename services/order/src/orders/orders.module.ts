import { Module } from '@nestjs/common';
import { OrdersService } from './orders.service';
import { OrdersController } from './orders.controller';
import { SequelizeModule } from '@nestjs/sequelize';
import { Order } from './entities/order.entity';
import { AccountsModule } from 'src/accounts/accounts.module';

@Module({
  imports: [
    AccountsModule,
    SequelizeModule.forFeature([
      Order
    ])
  ],
  controllers: [OrdersController],
  providers: [OrdersService]
})
export class OrdersModule {}
