/*
Copyright (c) 2020, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package commands

import "context"

type Pipeline interface {
	Process(ctx context.Context)
}
