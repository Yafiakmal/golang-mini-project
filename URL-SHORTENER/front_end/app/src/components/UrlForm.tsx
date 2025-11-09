import axios from 'axios'
import { useState } from "react"
import { useGlobal } from './GlobalState';
import { API_URL } from '../config/env';

export default function UrlForm() {
  const [url, setUrl] = useState("");
  const [urlShort, setUrlShort] = useState("");
  const { toggle, setToggle } = useGlobal();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await axios.post(`${API_URL}/shortener`, {
        url: url,
        name: urlShort,
      });

      console.log("Response:", response.data);
      setToggle(!toggle);
      alert("Success Response: " + JSON.stringify(response.data));
    } catch (error: any) {
      if (axios.isAxiosError(error)) {
        // Kalau error dari Axios (misalnya 400, 500, dll)
        console.error("Axios error:", error.response?.data || error.message);
        alert("Axios Error: " + error.message)
      } else {
        // Kalau error bukan dari Axios
        console.error("Unexpected error:", error);
        alert("Unexp Error: " + error.message)

      }
    } finally {
      setToggle(!toggle);
    }
  }
  return (
    <div className="flex items-center justify-center bg-gray-100">
      <form onSubmit={handleSubmit}
        className="bg-white p-8  rounded-2xl shadow-lg w-full max-w-md space-y-4">
        <h1 className="text-2xl font-semibold text-center text-gray-700">
          URL Shortener
        </h1>

        <input
          type="text"
          placeholder="Enter original URL"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          className="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
          required
        />

        <input
          type="text"
          placeholder="Enter short path (optional)"
          value={urlShort}
          onChange={(e) => setUrlShort(e.target.value)}
          className="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
        />

        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 rounded-lg transition"
        >
          Create Short URL
        </button>
      </form>
    </div>
  )

}

