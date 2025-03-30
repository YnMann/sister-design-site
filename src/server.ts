import express from 'express';
import dotenv from 'dotenv';
import cors from 'cors';
import path from 'path';
import morgan from 'morgan';
import htmlRoutes from './routes/html';

dotenv.config();

const app = express();
const PORT = process.env.PORT || 3000;

app.use(cors());
app.use(morgan('dev'));

// Middleware для парсинга JSON
app.use(express.json());

// Устанавливаем EJS как шаблонизатор
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views')); 

// Раздача статических файлов
app.use(express.static(path.join(__dirname, 'public')));

// Роуты
app.use('/', htmlRoutes);

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
