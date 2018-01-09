import React, { Component } from "react"

import Form from "../components/Form"

class Create extends Component {

    render(){
        return (
            <div className="container is-fluid">
                <p className="title is-3">Tracker</p>
                <p className="subtitle is-6">New tracker</p>

                <Form />
            </div>
        )
    }
}

export default Create
