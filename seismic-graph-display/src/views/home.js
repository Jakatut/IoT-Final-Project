import React, { useRef, useEffect } from "react"
import UnitLocation from "../components/UnitLocation";

const Home = () =>{

    const webSocket = useRef(null)
    const chartRef1 = useRef()
    const chartRef2 = useRef()


    useEffect(() => {
        webSocket.current = new WebSocket("ws://localhost:8000/readings")
        webSocket.current.onopen = () => console.log("ws opened")
        webSocket.current.onclose = () => console.log("ws closed")
        webSocket.current.onerror = (event) => {
            console.log("error: " + event)
        }

        if (!webSocket.current) {
            return
        }

        webSocket.current.onmessage = e => {
            const message = JSON.parse(e.data)
            console.log(message)
            if (message.name === "seismograph_1") {
                chartRef1.current.addData(message)
            } else if (message.name === "Attila1") {
                chartRef2.current.addData(message)
            }
        }
    })

    return(
        <div>            
            <UnitLocation Name="Barcelona" Latitude={41.3851} Longitude={2.1734} DeviceName={"sesimograph_1"} ref={chartRef1}/>
            <br/>
            <UnitLocation Name="Ocean" Latitude={7.772239} Longitude={-122.422889} DeviceName={"sesimograph_2"} ref={chartRef2}/>
        </div>
    )
}


export default Home