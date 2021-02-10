import React from "react"
import { GoogleMap, LoadScript } from '@react-google-maps/api';

const MapView = (props) =>{
    const latitude = props.Latitude
    const longitude = props.Longitude
    const apiKey = process.env.REACT_APP_GOOGLE_MAP_KEY

    // const latitude = 69.7
    // const longitude = 123.5


    const mapStyles = {        
        height: "20vh",
        width: "100%"};
      
      const defaultCenter = {
        lat: latitude, lng: longitude
      }

    return(
        <LoadScript
        googleMapsApiKey={apiKey}>
         <GoogleMap
           mapContainerStyle={mapStyles}
           zoom={13}
           center={defaultCenter}
         />
      </LoadScript>
    )
}

export default MapView
