import { Router, Request, Response } from 'express';

const router = Router();

// Роут для отдачи главной страницы (main.ejs)
router.get('/', (req: Request, res: Response) => {
  res.render('main', { title: 'Home Page', message: 'sad life' });
});

export default router;
