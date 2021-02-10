import React from "react"
import { VictoryChart, VictoryLine } from 'victory'

  
const LineGraph = (props) =>{
    return(
      <div>        
        <VictoryChart>
          <VictoryLine
            data={props.data}
            x="id"
            y="scale"      
          />
        </VictoryChart>
      </div>
    )
}

export default LineGraph