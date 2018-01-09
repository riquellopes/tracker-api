import React, { Component } from "react"

class Row extends Component {

    render(){
        return (
            <tr>
                <td>{this.props.id}</td>
                <td>{this.props.name}</td>
                <td>{this.props.data}</td>
                <td>{this.props.opened}</td>
                <td>{this.props.create_at}</td>
                <td>{this.props.update_at}</td>
            </tr>
        )
    }
}


export default Row
