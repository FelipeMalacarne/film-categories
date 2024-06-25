import { Link } from "react-router-dom";
import { cn } from "../lib/utils";
import { Package2 } from "lucide-react";
import { useState } from "react";

export const Navbar = () => {
    const [pathName, setPathName] = useState<string>(window.location.pathname);

    return (
        <header className="sticky top-0 flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6">
            <nav className="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
                <Link
                    to="#"
                    className="flex items-center gap-2 text-lg font-semibold md:text-base"
                >
                    <Package2 className="h-6 w-6" />
                    <span className="sr-only">Filmes</span>
                </Link>
                <Link
                    to="/films"
                    className={cn("transition-colors hover:text-foreground", pathName === "/films" ? "text-foreground" : "text-muted-foreground")}
                    onClick={() => setPathName("/films")}
                >
                    Films
                </Link>
                <Link
                    to="/categories"
                    className={cn("transition-colors hover:text-foreground", pathName === "/categories" ? "text-foreground" : "text-muted-foreground")}
                    onClick={() => setPathName("/categories")}
                >
                    Categories
                </Link>
            </nav>
        </header>
    );
};
