import { useState } from "react";
import axios from "axios";

export default function Home() {
  const domains = [
    "gmail.com",
    "yahoo.com",
    "outlook.com",
    "hotmail.com",
    "aol.com",
    "icloud.com",
    "protonmail.com",
    "zoho.com",
    "yandex.com",
    "mail.com",
    "gmx.com",
    "tutanota.com",
    "fastmail.com",
    "mail.ru",
    "inbox.lv",
    "seznam.cz",
    "bk.ru",
    "rambler.ru",
    "list.ru",
    "yandex.ru",
    "ya.ru",
    "narod.ru",
    "hotmail.co.uk",
    "yahoo.co.uk",
    "live.co.uk",
    "btinternet.com",
    "virginmedia.com",
    "blueyonder.co.uk",
    "ntlworld.com",
    "sky.com",
    "talktalk.net",
    "tiscali.co.uk",
  ];

  const [email, setEmail] = useState("");
  const [suggestedDomain, setSuggestedDomain] = useState([]);
  const [response, setResponse] = useState("");
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => {
    setEmail(e.target.value);
    if (e.target.value && e.target.value.includes("@")) {
      const domain = e.target.value.split("@")[1];
      if (domain && domain.length > 0) {
        const filteredDomains = domains.filter((d) => d.includes(domain));
        setSuggestedDomain(filteredDomains);
      }
    } else {
      setSuggestedDomain([]);
    }
  };

  const handleVerify = (e) => {
    e.preventDefault();

    setResponse("");
    setSuggestedDomain([]);
    setLoading(true);

    axios
      .get(import.meta.env.VITE_API_URL + "/api/verify?email=" + email)
      .then((res) => {
        if (res.data) {
          setResponse("Email is valid and deliverable.");
        }
      })
      .catch((err) => {
        setResponse(
          err.response?.data || "An error occurred. Please try again."
        );
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <div className="bg-gray-50 text-gray-800 min-h-screen flex flex-col">
      {/* Hero Section */}
      <header className="bg-beige-200 text-gray-900 py-20">
        <div className="container max-w-2xl mx-auto text-center px-6">
          <h1 className="text-4xl font-extrabold mb-4">
            Verify Email Addresses Instantly
          </h1>
          <p className="text-lg mb-10">
            Enhance your email lists with our reliable, fast email verification
            service.
          </p>
          <p
            className={`text-sm font-medium p-4 rounded-lg mt-4 transition-all duration-300 ${
              !response
                ? "hidden"
                : response === "Email is valid and deliverable."
                ? "bg-green-100 text-green-700 border border-green-300"
                : "bg-red-100 text-red-700 border border-red-300"
            }`}
          >
            {response}
          </p>
          <form
            onSubmit={handleVerify}
            className="mt-4 mx-auto flex flex-col sm:flex-row items-center space-y-4 sm:space-y-0"
          >
            <input
              type="email"
              name="email"
              value={email}
              onChange={handleChange}
              className="w-full px-4 py-3 text-gray-700 rounded-md border focus:outline-none focus:ring-2 focus:ring-green-400"
              placeholder="example@domain.com"
              required
            />
            <button
              type="submit"
              className={`w-full sm:w-auto bg-green-500 hover:bg-green-600 text-white font-semibold py-3 px-8 rounded-md transition duration-300 sm:ml-4 ${
                loading && "opacity-50 cursor-not-allowed"
              }`}
              disabled={loading}
            >
              {loading ? "Verifying..." : "Verify"}
            </button>
          </form>
          <div className="mt-4 w-full flex justify-center items-center flex-wrap gap-2">
            {suggestedDomain.map((domain) => (
              <button
                key={domain}
                onClick={() => {
                  setEmail((prev) => prev.split("@")[0] + "@" + domain);
                  setSuggestedDomain([]);
                }}
                className="text-sm text-gray-700 hover:text-white mr-3 py-2 px-4 rounded-full border border-gray-300 bg-gray-50 hover:bg-green-500 shadow-sm transition duration-200 ease-in-out transform hover:scale-105"
              >
                @{domain}
              </button>
            ))}
          </div>
        </div>
      </header>

      {/* Benefits Section */}
      <section className="py-16 bg-white">
        <div className="container mx-auto px-6 text-center">
          <h2 className="text-3xl font-bold mb-8 text-gray-900">
            Why Choose Our Verifier?
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="p-6 bg-gray-50 rounded-lg shadow-sm hover:shadow-md transition">
              <h3 className="text-xl font-semibold mb-4">Fast & Accurate</h3>
              <p>
                Quick and precise validation ensures your emails are always
                deliverable and clean.
              </p>
            </div>
            <div className="p-6 bg-gray-50 rounded-lg shadow-sm hover:shadow-md transition">
              <h3 className="text-xl font-semibold mb-4">
                Reduced Bounce Rate
              </h3>
              <p>
                Prevent bounce rates by filtering out invalid and risky
                addresses with ease.
              </p>
            </div>
            <div className="p-6 bg-gray-50 rounded-lg shadow-sm hover:shadow-md transition">
              <h3 className="text-xl font-semibold mb-4">
                User-Friendly & Free
              </h3>
              <p>
                Designed with simplicity in mind, allowing easy, no-cost email
                verification for everyone.
              </p>
            </div>
          </div>
        </div>
      </section>

      <section className="container p-8 rounded-lg mx-auto my-10">
        <h2 className="text-2xl text-center font-bold text-gray-800 mb-4">
          About This Project
        </h2>

        <p className="text-gray-700 mb-4">
          This project's backend is hosted on Render, which means the server may
          go into a "sleep" mode when not in use. As a result, there can be a
          spin-up time of around 2 minutes for the API to respond initially
          after inactivity. This setup allows us to provide free hosting without
          compromising functionality, though it may introduce a short delay for
          the first request.
        </p>

        <h3 className="text-xl text-center font-semibold text-gray-800 mt-6 mb-2">
          About the Developer
        </h3>

        <p className="text-gray-700 mb-4">
          Hi there! I'm a passionate individual developer dedicated to building
          and deploying innovative solutions. If you'd like to check out more of
          my work, feel free to explore my{" "}
          <a
            href="https://github.com/utkarshagrrawal"
            target="_blank"
            rel="noopener noreferrer"
            className="underline text-blue-500"
          >
            gitbub profile.
          </a>
        </p>
      </section>

      {/* Footer Section */}
      <footer className="bg-gray-900 text-gray-400 py-8">
        <div className="container mx-auto px-6 text-center">
          <p>
            &copy; {new Date().getFullYear()} EmailVerifier. All rights
            reserved.
          </p>
          <p className="mt-2">
            Built with precision for reliable email verification.
          </p>
        </div>
      </footer>
    </div>
  );
}
