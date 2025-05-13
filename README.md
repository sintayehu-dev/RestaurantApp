# 🍽️ RestaurantApp

A powerful, full-featured Restaurant Management System API built with Go, providing comprehensive functionality for restaurant operations.

![License](https://img.shields.io/badge/license-MIT-blue)
![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8)
![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791)

## 🌟 Features

- **User Authentication** - Secure JWT-based authentication with role-based access control
- **Menu Management** - Create, update, and organize menu items and categories
- **Table Management** - Track table availability and status
- **Order Processing** - Comprehensive order lifecycle management
- **Invoice Generation** - Generate and manage customer invoices
- **Role-Based Access** - Different permission levels for staff and administrators

## 📋 API Endpoints

### Authentication
- `POST /users/signup` - Register a new user
- `POST /users/login` - Authenticate and receive access tokens

### Users
- `GET /users` - List all users (admin only)
- `GET /users/:user_id` - Get specific user details

### Tables
- `GET /tables` - List all tables
- `GET /tables/:table_id` - Get specific table details
- `GET /available-tables` - List currently available tables
- `POST /tables` - Add a new table (admin only)
- `PATCH /tables/:table_id` - Update a table (admin only)
- `DELETE /tables/:table_id` - Remove a table (admin only)

### Menus
- `GET /menus` - List all menus
- `GET /menus/:menu_id` - Get specific menu details
- `GET /menu-categories` - List all menu categories
- `POST /menus` - Create a new menu (admin only)
- `PATCH /menus/:menu_id` - Update a menu (admin only)
- `DELETE /menus/:menu_id` - Remove a menu (admin only)

### Food Items
- `GET /foods` - List all food items
- `GET /foods/:food_id` - Get specific food item details
- `GET /foods/category/:category` - List food items by category
- `GET /foods/search` - Search food items
- `POST /foods` - Add a new food item (admin only)
- `PATCH /foods/:food_id` - Update a food item (admin only)
- `DELETE /foods/:food_id` - Remove a food item (admin only)

### Orders
- `GET /orders` - List all orders (admin only)
- `GET /orders/:order_id` - Get specific order details
- `GET /user/orders` - List current user's orders
- `POST /orders` - Create a new order
- `PATCH /orders/:order_id` - Update an order (admin only)
- `DELETE /orders/:order_id` - Remove an order (admin only)

### Order Items
- `GET /orderItems` - List all order items (admin only)
- `GET /orderItems/:order_item_id` - Get specific order item details
- `GET /orders/:order_id/items` - List items in an order
- `POST /orderItems` - Add an item to an order
- `PATCH /orderItems/:order_item_id` - Update an order item
- `DELETE /orderItems/:order_item_id` - Remove an item from an order

### Invoices
- `GET /invoices` - List all invoices (admin only)
- `GET /invoices/:invoice_id` - Get specific invoice details
- `GET /user-invoices` - List current user's invoices
- `POST /invoices` - Generate a new invoice (admin only)
- `PATCH /invoices/:invoice_id` - Update an invoice (admin only)
- `DELETE /invoices/:invoice_id` - Remove an invoice (admin only)

## 🔧 Installation & Setup

### Prerequisites
- Go 1.24 or higher
- PostgreSQL
- Git

### Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/RestaurantApp.git
   cd RestaurantApp
   ```

2. **Set up environment variables**
   Create a `.env` file in the root directory with the following variables:
   ```
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=restaurant_db
   DB_PORT=5432
   PORT=9000
   SECRET_KEY=your_secret_key
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

## 📝 Database Schema

The application uses the following models:
- `User` - Authentication and user management
- `Table` - Restaurant tables information
- `Menu` - Menu categories and organization
- `Food` - Food items with prices and details
- `Order` - Customer orders with status tracking
- `OrderItem` - Individual items within an order
- `Invoice` - Payment information for completed orders
- `Note` - Additional notes and information

## 🧪 Testing

Run the test suite with:

```bash
go test ./...
```

## 🚀 Deployment

The application can be deployed as a standalone API or as part of a larger system:

- **Docker**
  ```bash
  docker build -t restaurant-app .
  docker run -p 9000:9000 restaurant-app
  ```

- **Kubernetes** - Deployment templates available in the `/deployment` directory

## 🔒 Security Features

- JWT-based authentication
- Password hashing using bcrypt
- Role-based access control
- Request validation
- Environment-based configuration

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 📧 Contact

If you have any questions or feedback, please reach out to:

- Your Name - [your.email@example.com](mailto:your.email@example.com)
- Project Link: [https://github.com/yourusername/RestaurantApp](https://github.com/yourusername/RestaurantApp)

---

Made with ❤️ using Go, Gin, and GORM
