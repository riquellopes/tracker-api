import React, { Component } from "react"

// Components
// import Row from "./Row"

class Grid extends Component {

    render(){
        // const recordscalled = this.props.list;

        return (
            <table className="table is-bordered is-striped is-narrow is-fullwidth">
                <thead>
                    <tr>
                        <th>Id</th>
                        <th>Name</th>
                        <th>Data</th>
                        <th>Opened</th>
                        <th>CreateAt</th>
                        <th>UpdatedAt</th>
                    </tr>
                </thead>
                <tbody>
                    {/* {recordscalled.map((item, index) => <Row {...item} key={index} />)} */}
                </tbody>
            </table>
        )
    }
}

export default Grid
