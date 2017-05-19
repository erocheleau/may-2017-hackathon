export const typeDefs = `
  type Result {
    title: String
    uniquerID: String
  }

  type Query {
    results: [Results]
  }
`;