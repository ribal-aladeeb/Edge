import React, { Component } from "react";

class SearchGroup extends Component {
  constructor() {
    super();
    this.state = {
      search: "",
      groups: []
    };
  }

  componentDidMount() {
    fetch("http://localhost:8000/groups")
      .then(res => res.json())
      .then(json => {
        this.setState({
          groups: json
        });
      });
  }

  getList() {
    if (this.state.search !== "") {
      let filteredSearch = this.state.groups.filter(group => {
        //if you cannot find this search within it, do not return it
        return group.name.toLowerCase().indexOf(this.state.search) !== -1;
      });
      return (
        <ul>
          {filteredSearch.map(item => (
            <div className="cardmessage">
              <li key={item.id}>
                <h5>
                  <div className="title">
                    <a href={"/group/" + item.id}>{item.name}</a>
                  </div>
                </h5>
                <h10>
                  <div className="description">{item.description}</div>
                </h10>
              </li>
            </div>
          ))}
        </ul>
      );
    }
  }

  updateSearch(event) {
    this.setState({ search: event.target.value.substr(0, 20) });
  }

  render() {
    return (
      <div>
        <center>
          <h4>Group Search</h4>
          <input
            type="text"
            value={this.state.search}
            onChange={this.updateSearch.bind(this)}
          />
          {this.getList()}
        </center>
      </div>
    );
  }
}

export default SearchGroup;