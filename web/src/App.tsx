import { Outlet } from "react-router";
import { useHealthCheck } from "./api/api";

function App() {
  const { data, isPending, isError } = useHealthCheck();

  return (
    <>
      <h1>FSLedger</h1>
      API Status:{" "}
      {isPending ? "Loading" : isError || data.status != 204 ? "Error" : "OK"}
      <Outlet />
    </>
  );
}

export default App;
