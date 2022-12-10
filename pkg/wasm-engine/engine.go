package wasmengine

import (
	"io/ioutil"

	"github.com/bytecodealliance/wasmtime-go"
	"github.com/sirupsen/logrus"
)

type WASIRunner struct {
	engine *wasmtime.Engine
	linker *wasmtime.Linker
	log *logrus.Entry
}

func NewWASIRunner(log *logrus.Entry) (*WASIRunner, error) {
	log.Info("Creating new wasi runner")
	engine := wasmtime.NewEngine()
	wasi := &WASIRunner{
		engine: engine,
		linker: wasmtime.NewLinker(engine),
		log: log,
	}

	err := wasi.linker.DefineWasi()

	if err != nil {
		return nil, err
	}

	return wasi, nil
}

func (w *WASIRunner) GenerateNewModuleFromWASM(wasm []byte) (*wasmtime.Module, error) {
	w.log.Info("Compiling wasm")
	// compiling module from wasm
	return wasmtime.NewModule(w.engine,wasm)
}

func (w *WASIRunner) GenerateNewWASMStoreFromConfig(stdOutPath, stdErrorPath string, directoryMapings map[string]string, args []string ) (*wasmtime.Store) {
	// create new store
	store := wasmtime.NewStore(w.engine)

	// new wasi config
	config := wasmtime.NewWasiConfig()

	// setting the out and error files
	w.log.Debugf("Mapping stdout and stderr to %s & %s",stdOutPath,stdErrorPath)
	config.SetStdoutFile(stdOutPath)
	config.SetStderrFile(stdErrorPath)
	config.SetArgv(args)

	// preopening defined directory mappings
	for internal, external := range directoryMapings {
		w.log.Debugf("Mapping directories from %s => %s",internal,external)
		config.PreopenDir(internal,external)
	}

	// seting the config to store
	store.SetWasi(config)

	w.log.Info("New store with wasi config")

	return store
}

func (w *WASIRunner) Run(wasmPath, stdOutPath, stdErrorPath string, directoryMapings map[string]string, args []string ) error {
	// read wasm
	w.log.Info("Reading wasm")
	data , err := ioutil.ReadFile(wasmPath)
	if err != nil {
		w.log.Errorf("Unable to read wasm path %v",err)
		return err
	}

	// compile the wasm
	module, err := w.GenerateNewModuleFromWASM(data)
	if err != nil {
		w.log.Errorf("Unable to get module %v",err)
		return err
	}

	// create store with wasi config
	store := w.GenerateNewWASMStoreFromConfig(stdOutPath,stdErrorPath,directoryMapings,args)

	// linking them
	instance, err := w.linker.Instantiate(store, module)
	if err != nil {
		w.log.Errorf("Unable to instatiate linker %v",err)
		return err
	}

	// Get _start function
	start := instance.GetFunc(store, "_start",) 
	if start == nil {
		w.log.Errorf(ERR_GETTING_START.Error())
		return ERR_GETTING_START
	}

	w.log.Info("running start")
	_, err = start.Call(store)
	
	return err
}