import { createBrowserRouter } from "react-router";
import App from "./App.tsx";
import ErrorPage from "./pages/ErrorPage.tsx";
import Index from "./pages/Index.tsx";
import NotFound from "./pages/NotFound.tsx";
import PreferenceStocks from "./pages/preference/Stocks.tsx";
import PreferenceTags from "./pages/preference/Tags.tsx";
import TransactionNew from "./pages/transaction/New.tsx";
import TransactionEdit from "./pages/transaction/Edit.tsx";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: App,
    errorElement: <ErrorPage />,
    children: [
      { index: true, Component: Index, id: "index" },
      { path: "/transaction/new", Component: TransactionNew },
      { path: "/transaction/edit/:id", Component: TransactionEdit },
      { path: "/preference/stocks", Component: PreferenceStocks },
      {
        path: "/preference/tags/stock",
        element: <PreferenceTags tagType="stock" />,
      },
      {
        path: "/preference/tags/transaction",
        element: <PreferenceTags tagType="transaction" />,
      },
      {
        path: "/preference/tags/flow",
        element: <PreferenceTags tagType="flow" />,
      },
      { path: "*", Component: NotFound },
    ],
  },
]);
