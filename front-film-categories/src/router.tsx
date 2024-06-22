import App from "./App";
import { createBrowserRouter } from "react-router-dom";
import FilmPage from "./pages/film-page/page";
import SupplierPage from "./pages/supplier-page/page";
import CategoriesPage from "./pages/category-page/page";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "films",
        element: <FilmPage />,
      },
      {
        path: "categories",
        element: <CategoriesPage />,
      },
      {
        path: "suppliers",
        element: <SupplierPage />,
      },
    ],
  },
]);
