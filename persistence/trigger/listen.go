package trigger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func listenToTriggers(pool *pgxpool.Pool, callback func(uint, string, uint, string) error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		os.Exit(1)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), fmt.Sprintf("listen %s", TriggerChannelInsertOrUpdate))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error listening to chat channel:", err)
		os.Exit(1)
	}

	for {
		notification, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for notification:", err)
			os.Exit(1)
		}

		payloads := strings.Split(notification.Payload, ";")
		rowID, _ := strconv.ParseUint(payloads[1], 10, 64)
		go callback(uint(notification.PID), payloads[0], uint(rowID), payloads[2])
	}
}
