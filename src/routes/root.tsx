import { Outlet, Link } from "react-router-dom";

export default function Root() {
  return (
    <div className="mx-auto flex min-h-screen max-w-5xl flex-col px-5 py-2">
      <header>
        <nav>
          <Link
            to="/"
            className="text-xl font-black text-black dark:text-white"
          >
            dnsd
          </Link>
        </nav>
      </header>
      <main className="grow">
        <Outlet />
      </main>
      <footer>
        <p className="text-center text-sm text-neutral-500">
          &copy; 2023 Zac Wojcik
        </p>
      </footer>
    </div>
  );
}
