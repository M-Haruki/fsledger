import { createBrowserRouter } from "react-router";
import App from "./App.tsx";
import ErrorPage from "./pages/ErrorPage.tsx";
import Index from "./pages/Index.tsx";
import NotFound from "./pages/NotFound.tsx";
import PreferenceStocks from "./pages/preference/Stocks.tsx";
import PreferenceTags from "./pages/preference/Tags.tsx";
import TransactionNew from "./pages/transaction/New.tsx";
import TransactionEdit from "./pages/transaction/Edit.tsx";
import { TransactionView } from "./pages/view/Transaction.tsx";
import { FlowView } from "./pages/view/Flow.tsx";
import { StockView } from "./pages/view/Stock.tsx";
import { NewFlowView } from "./pages/view/NewFlow.tsx";
import { NewTransactionView } from "./pages/view/NewTransaction.tsx";
import { NewStockView } from "./pages/view/NewStock.tsx";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: App,
    errorElement: <ErrorPage />,
    children: [
      { index: true, Component: Index, id: "index" },
      { path: "/transaction/new", Component: TransactionNew },
      { path: "/transaction/edit/:id", Component: TransactionEdit },
      { path: "/view/stock/new", Component: NewStockView },
      { path: "/view/stock/:id", Component: StockView },
      { path: "/view/transaction/new", Component: NewTransactionView },
      { path: "/view/transaction/:id", Component: TransactionView },
      { path: "/view/flow/new", Component: NewFlowView },
      { path: "/view/flow/:id", Component: FlowView },
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
