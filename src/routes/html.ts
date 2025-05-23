import { Router, Request, Response } from "express";
import { projects } from "../domain/index";

const router = Router();

const PORT = process.env.PORT || 4000;
const isProd = process.env.NODE_ENV === "production";
const host = isProd ? process.env.HOST : `localhost:${PORT}`;

// Роут для отдачи главной страницы (main.ejs)
router.get("/", (req: Request, res: Response) => {
  res.render("main", { host, projects });
});

// Роут для отдачи конкретного проекта (project.ejs)
router.get("/project/:name", (req: Request, res: Response) => {
  const projectName = req.params["name"];
  const project = projects.find((p) => p.name === projectName);

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


export default router;
