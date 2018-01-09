import React, { Component } from "react"


class Form extends Component {

    constructor() {
        super();

        this.state = {
            name: "",
            display: "none"
        }

        this.handleChange = this.handleChange.bind(this);
        this.handlerSubmit = this.handlerSubmit.bind(this);
    }
    handlerSubmit(event) {
        event.preventDefault();

        console.info(this.state);

        fetch("http://localhost:5000/tracker", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(this.state)
            })
            .then((response) => {
                return response.json()
            })
            .then((json) => {
                this.beforeResponse(json);
            })
            .catch((error) => {
                console.error("Fail", error);
            });
    }

    beforeResponse(response){
        this.refs.form.reset();
        this.setState({
            name: "",
            display: "block"
        });
    }

    handleClick(event){
        this.setState({"display": "none"});
    }

    handleChange(event) {
        this.setState({name: event.target.value});
    }

    render(){
        return (
            <div>
                <div className="notification is-success" style={{display:this.state.display}}>
                  <button className="delete" onClick={this.handleClick.bind(this)}></button>
                  Tracker added.
                </div>
                <form onSubmit={this.handlerSubmit} ref="form">
                    <div className="field">
                        <label className="label">Name:</label>
                        <div className="control">
                            <input type="text"
                                   name="name"
                                   // value={this.state.name}
                                   onChange={this.handleChange}
                                   className="input"
                                   required
                                   placeholder="Give a name for the tracker."/>
                        </div>
                    </div>
                    <br />
                    <div className="field">
                        <div className="control">
                            <p className="control">
                                <input type="submit" value="save" className="button is-primary"/>
                            </p>
                        </div>
                    </div>
                </form>
            </div>
        )
    }
}

export default Form
