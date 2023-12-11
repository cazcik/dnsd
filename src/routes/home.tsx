import { useState } from "react";

import type { ApiResponse } from "../../types";

export default function Home() {
  const [domain, setDomain] = useState("");
  const [jsonData, setJsonData] = useState<ApiResponse>(undefined);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleSubmit = async (e: any) => {
    e.preventDefault();

    setIsLoading(true);

    try {
      const res = await fetch(`/api/lookup`, {
        method: "POST",
        body: JSON.stringify({
          domain: domain,
        }),
        headers: {
          "Content-Type": "application/json",
        },
      });
      const data: ApiResponse = await res.json();

      setJsonData(data);
    } catch (error) {
      console.error("Error fetching data", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex flex-col py-24">
      <div className="flex w-full flex-col items-center justify-center">
        <h1 className="text-center text-xl font-black text-black dark:text-white">
          simple, dns research
        </h1>
        <form onSubmit={handleSubmit} className="mt-6 flex items-center">
          <input
            name="domain"
            value={domain}
            onChange={(e) => setDomain(e.target.value)}
            placeholder="example.org"
            className="rounded-bl-md rounded-tl-md border border-neutral-50 bg-neutral-50 px-3 py-2 outline-none placeholder:text-neutral-500 dark:border-neutral-800 dark:bg-neutral-800"
          />
          <button
            type="submit"
            disabled={isLoading}
            className="rounded-br-md rounded-tr-md border border-neutral-200 bg-neutral-200 px-3 py-2 disabled:border-neutral-300 disabled:bg-neutral-300 dark:border-neutral-700 dark:bg-neutral-700 dark:disabled:border-neutral-950 disabled:dark:bg-neutral-950"
          >
            {isLoading ? (
              <div role="status" className="px-1">
                <svg
                  aria-hidden="true"
                  className="h-6 w-6 animate-spin fill-neutral-800 text-neutral-200 dark:fill-neutral-200 dark:text-neutral-800"
                  viewBox="0 0 100 101"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                    fill="currentColor"
                  />
                  <path
                    d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                    fill="currentFill"
                  />
                </svg>
                <span className="sr-only">Loading...</span>
              </div>
            ) : (
              "search"
            )}
          </button>
        </form>
      </div>
      {jsonData && jsonData.data?.nameservers != null ? (
        <div className="my-24 space-y-4">
          <h2 className="text-lg font-black text-black dark:text-white">
            {jsonData.data.domain}
          </h2>
          {jsonData.data.nameservers ? (
            <div className="flex flex-col overflow-x-auto">
              <h2 className="font-bold text-black dark:text-white">
                NS Records
              </h2>
              <table className="mt-4 w-full table-auto">
                <thead>
                  <tr className="border border-neutral-300 bg-white dark:border-neutral-700 dark:bg-black">
                    <th className="px-4 py-2 text-left text-xs uppercase text-neutral-500">
                      Host
                    </th>
                    <th className="px-4 py-2 text-left text-xs uppercase text-neutral-500">
                      IP
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-neutral-300 pt-4 dark:divide-neutral-700">
                  {jsonData.data.nameservers.map((ns) => (
                    <tr>
                      <td className="px-4 py-2 text-left text-sm">{ns.name}</td>
                      <td className="px-4 py-2 text-left text-sm">{ns.ip}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : null}
          {jsonData.data.mx ? (
            <div className="flex flex-col overflow-x-auto">
              <h2 className="font-bold text-black dark:text-white">
                MX Records
              </h2>
              <table className="mt-4 w-full table-auto">
                <thead>
                  <tr className="border border-neutral-300 bg-white dark:border-neutral-700 dark:bg-black">
                    <th className="px-4 py-2 text-left text-xs uppercase text-neutral-500">
                      Host
                    </th>
                    <th className="px-4 py-2 text-left text-xs uppercase text-neutral-500">
                      IP
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-neutral-300 pt-4 dark:divide-neutral-700">
                  {jsonData.data.mx.map((m) => (
                    <tr>
                      <td className="px-4 py-2 text-left text-sm">{m.name}</td>
                      <td className="px-4 py-2 text-left text-sm">{m.ip}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : null}
          {jsonData.data.txt ? (
            <div className="flex flex-col overflow-x-auto">
              <h2 className="font-bold text-black dark:text-white">
                TXT Records
              </h2>
              <table className="mt-4 w-full table-auto">
                <thead>
                  <tr className="border border-neutral-300 bg-white dark:border-neutral-700 dark:bg-black">
                    <th className="px-4 py-2 text-left text-xs uppercase text-neutral-500">
                      Value
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-neutral-300 pt-4 dark:divide-neutral-700">
                  {jsonData.data.txt.map((t) => (
                    <tr>
                      <td className="px-4 py-2 text-left text-sm">{t}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : null}
          {jsonData.data.host ? (
            <div className="flex flex-col overflow-x-auto">
              <h2 className="font-bold text-black dark:text-white">
                A Records
              </h2>
              <table className="mt-4 w-full table-auto">
                <thead>
                  <tr className="border border-neutral-300 bg-white dark:border-neutral-700 dark:bg-black">
                    <th className="px-4 py-2 text-left text-xs uppercase text-neutral-500">
                      Host
                    </th>
                    <th className="px-4 py-2 text-left text-xs uppercase text-neutral-500">
                      IP
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-neutral-300 pt-4 dark:divide-neutral-700">
                  {jsonData.data.host.map((h) => (
                    <tr>
                      <td className="px-4 py-2 text-left text-sm">{h.name}</td>
                      <td className="px-4 py-2 text-left text-sm">{h.ip}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : null}
        </div>
      ) : (
        <p className="mt-10 text-center text-neutral-500">
          this domain is either unregisitered or contains an invalid tld
        </p>
      )}
    </div>
  );
}
