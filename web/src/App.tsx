import { Outlet } from "react-router";
import { useHealthCheck } from "./api/api";
import SideBar from "./components/SideBar";

function App() {
  const { data, isPending, isError } = useHealthCheck();

  return (
    <div className="flex bg-background">
      <SideBar className="w-3xs h-screen" />
      <h1>FSLedger</h1>
      API Status:{" "}
      {isPending ? "Loading" : isError || data.status != 204 ? "Error" : "OK"}
      <Outlet />
    </div>
  );
}

export default App;
