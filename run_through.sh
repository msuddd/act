echo Discovery, what happends on a push 
tree . 
act -g push
act -l push
read
echo Discovery, what happends on a workflow_dispatch 
act -g workflow_dispatch
act -l workflow_dispatch
read
echo "Dry Run"
act -n
read
