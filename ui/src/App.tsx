import { ThemeContextProvider } from "./context/theme"
import AppRouter from "./router"

export default function App(): JSX.Element {
  return (
    <div>
      <ThemeContextProvider>
        <AppRouter />
      </ThemeContextProvider>
    </div>
  )
}
