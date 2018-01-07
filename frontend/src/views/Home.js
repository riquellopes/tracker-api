import React, { Component } from "react"

// Components
import Grid from "../components/Grid"

class Home extends Component {

    constructor(){
      super();
      this.state = { trackers: [] };
    }

    componentWillMount(){
      fetch("http://localhost:5000/tracker")
        .then((response) => {
            return response.json()
        })
        .then((trackers) => {
            this.setState( {trackers: trackers})
        });
    }

    render(){
        return (
            <div className="container is-fluid">
                <p className="title is-3">Trackers</p>
                <Grid list={this.state.trackers}/>
            </div>
        )
    }
}

export default Home
