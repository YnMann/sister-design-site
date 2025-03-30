"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
const index_1 = require("../domain/index");
const router = (0, express_1.Router)();
const PORT = process.env.PORT || 4000;
const isProd = process.env.NODE_ENV === "production";
const host = isProd ? process.env.HOST : `localhost:${PORT}`;
// Роут для отдачи главной страницы (main.ejs)
router.get("/", (req, res) => {
    res.render("main", { host, projects: index_1.projects });
});
// Роут для отдачи конкретного проекта (project.ejs)
router.get("/project/:name", (req, res) => {
    const projectName = req.params["name"];
    const project = index_1.projects.find((p) => p.name === projectName);
    if (!project) {
        res.render("404", { host });
        return;
    }
    res.render("project", { host, project });
});
setInterval(() => {
    const used = process.memoryUsage();
    console.log(`Memory Usage: RSS=${(used.rss / 1024 / 1024).toFixed(2)} MB, Heap=${(used.heapUsed / 1024 / 1024).toFixed(2)} MB`);
}, 5000);
exports.default = router;
