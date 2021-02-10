import React, { useState, forwardRef, useImperativeHandle } from "react"
import LineGraph from "../components/LineGraph"
import MapView from "../components/map"
import Card from '@material-ui/core/Card'
import CardHeader from '@material-ui/core/CardHeader'
import CardContent from '@material-ui/core/CardContent'


function getRandomInt() {

    // 2 decimals
    var precision = 100

    var randomnum = (Math.floor(Math.random() * (10 * precision - 1 * precision) + 1 * precision) / (1*precision) / 10)
    return randomnum
  }


function checkLevel(level) {
    const danger = process.env.REACT_APP_THRESHOLD_DANGER
    const warning = process.env.REACT_APP_THRESHOLD_WARNING
    
    if (level > danger)
        return getLevel(1)
    else if (level > warning)
        return getLevel(2)
    else 
        return getLevel(0)

}

function getCurrentDate(){
    var d = new Date()
    var datestring = d.getDate()  + "-" + (d.getMonth()+1) + "-" + d.getFullYear() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds()
    return datestring

}

function getLevel(level) {
    const styleName = "list-group-item list-group-item-"

      // DANGER
      if (level === 1){
          return styleName + 'danger'

      }

      // WARNING
      else if (level === 2){
          return styleName + 'warning'
      }
      else{
          return styleName + 'light'
      }
      
  }

const UnitLocation = (props, ref) =>{        
    const [chartData, setChartData] = useState([])  
    const [count, setCount] = useState(0);
    const [timestamp, setTimestamp] = useState("")

    useImperativeHandle(ref, () => ({
        addData(data) {        
            setTimestamp(getCurrentDate)
            setCount(count + 1)

            const newData = {
                date: data.time,
                scale: data.scale
            }
            setChartData(items => [...items, newData])
        }
    }));

    return(
        <div className="container">

            <Card>                
      <CardHeader
        title={props.Name}
        subheader={timestamp}
      />
      <CardContent>
      <div className="row">
                <div className="col-4">
                <MapView Latitude={props.Latitude} Longitude={props.Longitude} />                    
                </div>
                <div className="col-8">
                <LineGraph data={chartData} />
                </div>
            </div>            
            <div className="row">
                <div className="col-12">
                <div style={{maxHeight: "200px", overflowY: 'scroll'}}>
            <ul class="list-group">            
                {chartData.map((x) => (
                    <li className={checkLevel(x.scale)} key={x.id}>
                        id: {x.id} scale: {x.scale}
                    </li>
                ))}
            </ul>
            </div> 
                </div> 
            </div>
    </CardContent>
    </Card>

            
    </div>
    )
}


export default forwardRef(UnitLocation)