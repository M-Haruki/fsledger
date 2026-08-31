import { Outlet } from "react-router";
// import { useHealthCheck } from "./api/api";
import SideBar from "./components/SideBar";

function App() {
  // const { data, isPending, isError } = useHealthCheck();

  return (
    <div className="flex bg-background w-screen h-screen">
      <SideBar className="w-3xs h-full p-3" />
      {/* API Status:{" "}
      {isPending ? "Loading" : isError || data.status != 204 ? "Error" : "OK"} */}
      <div className="flex-1 h-full p-3 relative">
        <Outlet />
      </div>
    </div>
  );
}

export default App;
