"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.projects = void 0;
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const project_url = "projects";
function getProjectImages(projectName) {
    const projectPath = path_1.default.join(__dirname, "../public", "projects", projectName);
    try {
        const files = fs_1.default.readdirSync(projectPath);
        return files
            .filter(file => /\.(png|jpg|jpeg|webp)$/i.test(file))
            .map(file => `/projects/${projectName}/${file}`);
    }
    catch (error) {
        return [];
    }
}
exports.projects = [
    {
        name: "arman",
        title: "Проект Arman",
        year: "2024",
        style: "Современный с классическими элементами",
        description: "Жилой комплекс премиум-класса, сочетающий современные архитектурные решения с классическими деталями. Просторные квартиры, панорамные окна и высококачественные материалы отделки.",
        images: getProjectImages("arman"),
    },
    {
        name: "first_arena_park",
        title: "Arena Park – Первая очередь",
        year: "2023",
        style: "Современный",
        description: "Современный жилой комплекс с развитой инфраструктурой. Удобные планировки, закрытый двор без машин, детские и спортивные площадки.",
        images: getProjectImages("arena-park-1"),
    },
    {
        name: "second_arena_park",
        title: "Arena Park – Вторая очередь",
        year: "2022",
        style: "Минимализм",
        description: "Вторая очередь жилого комплекса с улучшенными планировками и просторными общественными зонами. Минималистичный дизайн, функциональность и уют.",
        images: getProjectImages("arena-park-2"),
    },
    {
        name: "belle_view",
        title: "Belle View Residence",
        year: "2022",
        style: "Современная классика",
        description: "Элитный жилой комплекс с видом на реку. Закрытая территория, подземный паркинг, просторные террасы и панорамные окна.",
        images: getProjectImages("belle-view"),
    },
    {
        name: "office_tower",
        title: "Office Tower",
        year: "2022",
        style: "Минимализм",
        description: "Современный бизнес-центр с панорамными видами на город. Просторные офисные помещения, развитая инфраструктура и инновационные инженерные решения.",
        images: getProjectImages("office"),
    },
];
