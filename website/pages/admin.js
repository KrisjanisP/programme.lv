import NavBar from '../components/navbar'
import Link from "next/link";
import {useState} from 'react'

async function deleteTask(code, apiURL) {
    const fetchOptions = {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: JSON.stringify({code: code}),
    };

    const response = await fetch(`${apiURL}/tasks/delete/` + code, fetchOptions);

    if (!response.ok) {
        const errorMessage = await response.text();
        console.error(errorMessage)
        return
    }

    return response.json()
}

async function createTaskSubmit(form_event) {
    form_event.preventDefault()
    const form = form_event.currentTarget;
    const url = form.action;

    const formData = new FormData(form);

    const response = await fetch(url, {method: "POST", body: formData});

    if (!response.ok) {
        const errorMessage = await response.text();
        console.error(errorMessage)
        return
    }

    const responseData = response.text();
    console.log(responseData)
    return responseData
}

export default function Admin(props) {

    const [tasks, setTasks] = useState(props.tasks)

    let refreshTable = async () => {
        const res = await fetch(`${props.apiURL}/tasks/list`)
        const tasks = await res.json()
        setTasks(tasks)
    }
    let deleteTaskAndRefreshTable = async (code) => {
        await deleteTask(code,props.apiURL);
        await refreshTable();
    }
    let displayTaskDeleteModal = async (name, code) => {
        document.getElementById("delete-task-modal-header").innerHTML = name
        document.getElementById("delete-task-modal-confirm").onclick = async () => {
            await deleteTaskAndRefreshTable(code);
            document.getElementById("delete-task-modal-close").click()
        }
        document.getElementById("delete-task-modal-show").click()
    }
    let createTaskSubmitAndRefresh = async (form_event) => {
        await createTaskSubmit(form_event)
        await refreshTable()
    }

    let admin_table_entries = []
    let TagList = (props)=> {
        let tag_entries = []
        let tags = props.tags;
        for(let tag of tags) {
            let bg = "bg-secondary"
            if(tag==="ProblemCon++") bg = "bg-primary"
            tag_entries.push(<span className={`badge ${bg} m-1`} key={tag}>{tag}</span>)
        }
        return (
            <>
                {tag_entries}
            </>
        )
    }
    if(props.tasks) {
        tasks.forEach((task) => {
            admin_table_entries.push(
                <tr key={task["code"]}>
                    <th scope="row">
                        <Link href={"/tasks/" + task["code"]}>
                            <a className="nav-link">{task["code"]}</a>
                        </Link>
                    </th>
                    <td>
                        <Link href={"/tasks/" + task["code"]}>
                            <a className="nav-link">{task["name"]}</a>
                        </Link>
                    </td>
                    <td>{task["version"]}</td>
                    <td><TagList tags={task["tags"]}/></td>
                    <td><span className="badge bg-danger">6.9</span></td>
                    <td>2</td>
                    <td>13</td>
                    <td>
                        <button type="button" className="btn btn-sm btn-primary me-1">Redi????t</button>
                        <button type="button" className="btn btn-sm btn-danger ms-1"
                                onClick={() => displayTaskDeleteModal(task["name"], task["code"])}>Izdz??st
                        </button>
                    </td>
                </tr>
            )
        })
    }


    let ErrorAlert = ({ msg }) => {
        if (msg) return (
            <div className="alert alert-danger text-center" role="alert">
                {msg}
            </div>)
        else return <></>
    }

    return (
        <div>
            <NavBar active_page={"admin"}/>
            <main className="container">
                <h1 className="my-4 text-center">administr??cija</h1>

                <form action={`${props.apiURL}/tasks/create`} onSubmit={createTaskSubmitAndRefresh}>
                    <div className="row">
                        <div className="mb-3 col">
                            <input className="form-control" type="file" name="task-file" accept={".zip"}/>
                        </div>
                        <div className={"col"}>
                            <button type="submit" className="btn btn-success">pievienot uzdevumu</button>
                        </div>
                    </div>
                </form>

                <ErrorAlert msg={props.error}/>
                <table className="table table-hover" style={{tableLayout: "fixed"}}>
                    <thead>
                    <tr>
                        <th scope="col">kods</th>
                        <th scope="col">nosaukums</th>
                        <th scope={"col"}>versija</th>
                        <th scope="col">birkas</th>
                        <th scope="col">gr??t??ba</th>
                        <th scope="col">atrisin??jumi</th>
                        <th scope="col">ies??t??jumi</th>
                        <th scope={"col"}>darb??bas</th>
                    </tr>
                    </thead>
                    <tbody>
                    {admin_table_entries}
                    </tbody>
                </table>

            </main>

            <div className="modal fade" id="delete-task-modal" tabIndex="-1">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h5 className="modal-title" id="delete-task-modal-header"></h5>
                            <button id="delete-task-modal-close" type="button" className="btn-close"
                                    data-bs-dismiss="modal"></button>
                        </div>
                        <div className="modal-body">
                            Vai esat p??rliecin??ti, ka v??laties dz??st ??o uzdevumu?
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">aizv??rt
                            </button>
                            <button type="button" className="btn btn-danger" id="delete-task-modal-confirm">dz??st
                                uzdevumu
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            <button type="button" className="btn btn-primary d-none" id="delete-task-modal-show" data-bs-toggle="modal"
                    data-bs-target="#delete-task-modal">
            </button>

        </div>
    )
}


// This gets called on every request
export async function getServerSideProps() {
    let result = {
        props: {
            apiURL: process.env.API_URL
        }
    }

    try {
        const res = await fetch(`${process.env.API_URL}/tasks/list`)
        result.props.tasks = await res.json()
    } catch (err) {
        result.props.error = "failed to fetch tasks from the API :("
    }

    return result
}
