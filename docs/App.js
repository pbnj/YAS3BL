/* global fetch, semanticUIReact, ReactDOM, React */
const { Header, Container, Segment, Table, Dropdown } = semanticUIReact

const YAS3BL =
  'https://raw.githubusercontent.com/petermbenjamin/YAS3BL/master/yas3bl.json'
class App extends React.Component {
  state = { leaks: [] }

  componentDidMount() {
    fetch(YAS3BL)
      .then(response => response.json())
      .then(data => {
        data.sort((a, b) => a.organization.localeCompare(b.organization))
        this.setState({ leaks: data })
      })
      .catch(err => console.error(err))
  }

  render() {
    const { leaks } = this.state
    return (
      <Container>
        <Segment.Group horizontal>
          <Segment>
            <Header as="h1">Yet Another S3 Bucket Leak</Header>
          </Segment>
        </Segment.Group>
        <Table celled>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Organization</Table.HeaderCell>
              <Table.HeaderCell>Count (Records Exposed)</Table.HeaderCell>
              <Table.HeaderCell>Data Exposed</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {leaks.map((l, i) => (
              <Table.Row key={i}>
                <Table.Cell>
                  <a href={l.url}>{l.organization}</a>
                </Table.Cell>
                <Table.Cell>{l.count}</Table.Cell>
                <Table.Cell>{l.data}</Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Container>
    )
  }
}

const root = document.getElementById('root')
const elem = <App />
ReactDOM.render(elem, root)
