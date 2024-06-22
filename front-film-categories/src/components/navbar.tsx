import { Link } from "react-router-dom";
import { cn } from "../lib/utils";
import { buttonVariants } from "./ui/button";

export const Navbar = () => {
  return (
    <nav className="border-b">
      <div className="flex items-center h-16 px-4 space-x-8">
        <div>
          {/* <img src={logo} alt="logo" /> */}
          <h1>Caralho Films</h1>
        </div>
        <div>
          <ul className="flex items-center space-x-4 ml-auto">
            <li>
              <Link
                className={cn(buttonVariants({ variant: "ghost" }))}
                to="/suppliers"
              >
                Suppliers
              </Link>
            </li>
            <li>
              <Link
                className={cn(buttonVariants({ variant: "ghost" }))}
                to="/films"
              >
                Films
              </Link>
            </li>
            <li>
              <Link
                className={cn(buttonVariants({ variant: "ghost" }))}
                to="/categories"
              >
                Suppliers
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};
