import { Layers, TableOfContents, Tag, Wrench } from "lucide-react";
import type React from "react";
import { useState } from "react";
import { Link, NavLink } from "react-router";

export default function SideBar({ className }: { className: string }) {
  return (
    <div className={"flex flex-col gap-y-3 bg-primary-lightest " + className}>
      <Link to="/" className="font-extrabold text-3xl mb-5 flex items-center">
        <img src="/favicon.ico" className="h-8 mr-2" />
        FSLedger
      </Link>
      <MenuLink
        url="/transaction/new"
        className="font-semibold border-primary-light border-2"
      >
        New Transaction
      </MenuLink>
      <Menu
        label={
          <>
            <TableOfContents className="inline align-sub" /> View
          </>
        }
      >
        <></>
      </Menu>
      <Menu
        label={
          <>
            <Wrench className="inline align-sub" /> Preference
          </>
        }
      >
        <MenuLink url="/preference/stocks">
          <Layers className="inline align-sub" size={22} /> Stocks
        </MenuLink>
        <MenuLink url="/preference/tags/stock">
          <Tag className="inline align-sub" size={22} /> Stock Tags
        </MenuLink>
        <MenuLink url="/preference/tags/transaction">
          <Tag className="inline align-sub" size={22} /> Transaction Tags
        </MenuLink>
        <MenuLink url="/preference/tags/flow">
          <Tag className="inline align-sub" size={22} /> Flow Tags
        </MenuLink>
      </Menu>
    </div>
  );
}

function MenuLink({
  children,
  url,
  className = "",
}: {
  children: React.ReactNode;
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
      {children}
    </NavLink>
  );
}

function Menu({
  label,
  children,
}: {
  label: React.ReactNode;
  children: React.ReactNode;
}) {
  const [isOpen, setIsOpen] = useState(true);
  return (
    <div className="flex flex-col">
      <div
        className="block rounded-lg p-2 font-semibold hover:bg-primary-lighter text-lg cursor-pointer"
        onClick={() => setIsOpen(!isOpen)}
      >
        {label}
      </div>
      <div className={`w-9/10 ml-auto + ${isOpen ? "block" : "hidden"}`}>
        {children}
      </div>
    </div>
  );
}
