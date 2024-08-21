import React from "react"
import {createRoot} from "react-dom/client"
import { Rules } from "./pages/Rules"
import "./App.css"


const root = createRoot(document.getElementById("root"))


root.render(<div className="container"><Rules /></div>)

