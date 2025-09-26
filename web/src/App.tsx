import { useEffect, useState } from "react";

type IdentityResponse = {
  machineIdentifier: string;
  version?: string;
  friendlyName?: string;
};

const App = () => {
  const [identity, setIdentity] = useState<IdentityResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch("/api/identity")
      .then(async (response) => {
        if (!response.ok) {
          throw new Error(`Request failed with status ${response.status}`);
        }
        return response.json();
      })
      .then(setIdentity)
      .catch((err: Error) => setError(err.message));
  }, []);

  return (
    <div className="min-h-screen bg-slate-950 text-slate-100">
      <header className="border-b border-slate-800 bg-slate-900/50">
        <div className="mx-auto flex max-w-5xl items-center justify-between px-6 py-4">
          <span className="text-xl font-semibold">plex-discovery</span>
          <span className="text-sm text-slate-400">Preview</span>
        </div>
      </header>
      <main className="mx-auto flex max-w-5xl flex-col gap-6 px-6 py-10">
        <section className="rounded-lg border border-slate-800 bg-slate-900/40 p-6 shadow-lg">
          <h1 className="text-2xl font-bold text-amber-400">Server identity</h1>
          <p className="mt-2 text-sm text-slate-400">
            This panel demonstrates the generated Plex API client by calling a placeholder endpoint.
          </p>
          {error && <p className="mt-4 text-sm text-red-400">{error}</p>}
          {identity ? (
            <dl className="mt-6 grid gap-3 sm:grid-cols-3">
              <div>
                <dt className="text-xs uppercase tracking-wide text-slate-500">Machine ID</dt>
                <dd className="text-lg font-mono text-slate-200">{identity.machineIdentifier}</dd>
              </div>
              <div>
                <dt className="text-xs uppercase tracking-wide text-slate-500">Version</dt>
                <dd className="text-lg text-slate-200">{identity.version ?? "Unknown"}</dd>
              </div>
              <div>
                <dt className="text-xs uppercase tracking-wide text-slate-500">Name</dt>
                <dd className="text-lg text-slate-200">{identity.friendlyName ?? "n/a"}</dd>
              </div>
            </dl>
          ) : (
            !error && <p className="mt-4 text-sm text-slate-300">Loading identity...</p>
          )}
        </section>
      </main>
    </div>
  );
};

export default App;
