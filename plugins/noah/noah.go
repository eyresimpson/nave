package noah

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"nave/plugins/noah/shared"
	"os"
)

// GreeterHello Here is a real implementation of Greeter
type GreeterHello struct {
	logger hclog.Logger
}

func (g *GreeterHello) Greet() string {
	g.logger.Debug("message from GreeterHello.Greet")
	//Create(g)
	//
	return "Hello!"
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "noahjones",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	greeter := &GreeterHello{
		logger: logger,
	}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"greeter": &shared.GreeterPlugin{Impl: greeter},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

//
//const (
//	dbDriver = "mysql"
//	dbDSN    = "root:root@tcp(localhost:3306)/zerocloud_db"
//	poolSize = 10
//)
//
//func Create(f *GreeterHello) {
//	db, err := initDB()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Perform database operations using the connection pool
//	for i := 1; i <= 20; i++ {
//		go func(id int) {
//			conn := dbPool.Get().(*sql.DB)
//			defer dbPool.Put(conn)
//
//			// Simulate a database query
//			query := fmt.Sprintf("SELECT * FROM zc_sys_ou")
//			rows, err := conn.Query(query)
//			if err != nil {
//				log.Println(err)
//				return
//			}
//			defer rows.Close()
//
//			// Process the query result (just printing in this example)
//			for rows.Next() {
//				var userID int
//				var username string
//				if err := rows.Scan(&userID, &username); err != nil {
//					log.Println(err)
//					return
//				}
//				fmt.Printf("User ID: %d, Username: %s\n", userID, username)
//				f.logger.Info("User ID: %d, Username: %s\n", userID, username)
//			}
//		}(i)
//	}
//
//	// Wait for all goroutines to finish
//	wg.Wait()
//}
//
//var (
//	dbPool *sync.Pool
//	wg     sync.WaitGroup
//)
//
//func initDB() (*sql.DB, error) {
//	db, err := sql.Open(dbDriver, dbDSN)
//	if err != nil {
//		return nil, err
//	}
//
//	// Initialize the connection pool
//	dbPool = &sync.Pool{
//		New: func() interface{} {
//			conn, _ := sql.Open(dbDriver, dbDSN)
//			return conn
//		},
//	}
//
//	// Pre-warm the connection pool
//	for i := 0; i < poolSize; i++ {
//		conn := dbPool.Get().(*sql.DB)
//		dbPool.Put(conn)
//	}
//
//	return db, nil
//}
//
//// 基础的数据库操作，多数属于底层的操作
//
//// 初始化操作（此链接将被保存在流中）
//func init() {
//
//}
//
//// 执行SQL
//func ExecSQL(dbType string) {
//
//}
//
//// 销毁数据库链接池
