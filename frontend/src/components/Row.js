import React, { Component } from "react"

class Row extends Component {

    render(){
        return (
            <tr>
                <td>{this.props.id}</td>
                <td>{this.props.name.name}</td>
                <td>{this.props.data.data}</td>
                <td>{this.props.opened.opened}</td>
                <td>{this.props.createat}</td>
                <td>{this.props.updatedAt}</td>
            </tr>
        )
    }
}
