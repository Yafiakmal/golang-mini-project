import { GlobalProvider } from "./components/GlobalState.tsx"
import UrlForm from "./components/UrlForm.tsx"
import UrlList from "./components/UrlList.tsx"
export default function App() {
  return (
    <GlobalProvider>
      <UrlForm />
      <UrlList />
    </GlobalProvider>
  )
}
