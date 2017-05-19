import React, { Component } from 'react';
import { gql, graphql } from 'react-apollo';

const ResultList = ({ data: { loading, error, queryResults } }) => {
  if (loading) {
    return <p> Loading ... </p>
  }
  if (error) {
    return <p> {error.message} </p>
  }

  return (
    <ul> {
      queryResults.results.map(res => <li key={res.UniqueId}>{res.title}</li>)
    }
    </ul>
  )
}

const resultListQuery = gql`{
queryResults(q: "event") {
    results {
      title
      UniqueId
    }
  }
}`;

export const ResultListWithData = graphql(resultListQuery)
  (ResultList)
// export default Querybox;