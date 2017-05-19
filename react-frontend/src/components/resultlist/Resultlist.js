import React, { Component } from 'react';
import { gql, graphql } from 'react-apollo';

import './Resultlist.css';

const ResultList = ({ data: { loading, error, queryResults } }) => {
  if (loading) {
    return <p> Loading ... </p>
  }
  if (error) {
    return <p> {error.message} </p>
  }

  return (
    <div className="result-list-container"> {
      queryResults.results.map(res => 
        <div className="result-container">
          <a href={res.clickableuri}>{res.title}</a><br/>
          <span>{res.Excerpt}</span>
        </div>)
    }
    </div>
  )
}

const resultListQuery = gql`{
queryResults(q: "event") {
    results {
      title
      clickableuri
      Excerpt
    }
  }
}`;

export const ResultListWithData = graphql(resultListQuery)
  (ResultList)
// export default Querybox;