import React, { Component } from 'react';
import TextField from 'material-ui/TextField';
import FontIcon from 'material-ui/FontIcon';

import './Querybox.css';

const iconStyles = {
  lineHeight: "48px",
};

export class Querybox extends Component {
  
  render() {
    return (
      <div className="querybox-container">
        <div className="querybox-input-container">
          <TextField
            hintText="Search"
            fullWidth={true}
            className="querybox-input"
          />
        </div>
        <div className="querybox-search-icon">
          <FontIcon className="material-icons" style={iconStyles}>search</FontIcon>
        </div>
      </div>
    )
  }
}

// export default Querybox;