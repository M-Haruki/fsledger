import type React from "react";
import { useState } from "react";
import { Link, NavLink } from "react-router";

export default function SideBar({ className }: { className: string }) {
  return (
    <div
      className={"flex flex-col p-3 gap-y-3 bg-primary-lightest " + className}
    >
      <Link to="/" className="font-extrabold text-3xl mb-5 flex items-center">
        <img src="/favicon.ico" className="h-8 mr-2" />
        FSLedger
      </Link>
      <MenuLink
        name="New Transaction"
        url="/transaction/new"
        className="font-semibold border-primary-light border-2"
      ></MenuLink>
      <Menu name="View">
        <></>
      </Menu>
      <Menu name="Preference">
        <MenuLink name="Stocks" url="/preference/stocks"></MenuLink>
        <MenuLink name="Stock Tags" url="/preference/tags/stock"></MenuLink>
        <MenuLink
          name="Transaction Tags"
          url="/preference/tags/transaction"
        ></MenuLink>
        <MenuLink name="Flow Tags" url="/preference/tags/flow"></MenuLink>
      </Menu>
    </div>
  );
}

function MenuLink({
  name,
  url,
  className = "",
}: {
  name: string;
  url: string;
  className?: string;
}) {
  return (
    <NavLink
      to={url}
      className={({ isActive }) =>
        `block rounded-lg p-2 text-lg ${className} ${isActive ? "bg-primary-light" : "hover:bg-primary-lighter"}`
      }
    >
      {name}
    </NavLink>
  );
}

function Menu({ name, children }: { name: string; children: React.ReactNode }) {
  const [isOpen, setIsOpen] = useState(true);
  return (
    <div className="flex flex-col">
      <div
        className="block rounded-lg p-2 font-semibold hover:bg-primary-lighter text-lg cursor-pointer"
        onClick={() => setIsOpen(!isOpen)}
      >
        {name}
      </div>
      <div className={`w-9/10 ml-auto + ${isOpen ? "block" : "hidden"}`}>
        {children}
      </div>
    </div>
  );
}
