"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const dotenv_1 = __importDefault(require("dotenv"));
const cors_1 = __importDefault(require("cors"));
const path_1 = __importDefault(require("path"));
const morgan_1 = __importDefault(require("morgan"));
const html_1 = __importDefault(require("./routes/html"));
dotenv_1.default.config();
const app = (0, express_1.default)();
const PORT = process.env.PORT || 3000;
app.use((0, cors_1.default)());
app.use((0, morgan_1.default)('dev'));
// Middleware для парсинга JSON
app.use(express_1.default.json());
// Устанавливаем EJS как шаблонизатор
app.set('view engine', 'ejs');
app.set('views', path_1.default.join(__dirname, 'views'));
// Раздача статических файлов
app.use(express_1.default.static(path_1.default.join(__dirname, 'public')));
// Роуты
app.use('/', html_1.default);
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
