import { createBrowserRouter } from "react-router";
import App from "./App.tsx";
import ErrorPage from "./pages/ErrorPage.tsx";
import Index from "./pages/Index.tsx";
import NotFound from "./pages/NotFound.tsx";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: App,
    errorElement: <ErrorPage />,
    children: [
      { index: true, Component: Index, id: "index" },
      { path: "*", Component: NotFound },
    ],
  },
]);
