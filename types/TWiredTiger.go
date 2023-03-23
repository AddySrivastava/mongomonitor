package types

type TWiredTiger struct {
	Metadata struct {
		FormatVersion int32
	}
	CreationString string
	Type           string
	Uri            string
	Block_Manager  struct {
		Allocations_Requiring_File_Extension int32 `bson:"allocations requiring file extension"`
		Blocks_Allocated                     int32 `bson:"blocks allocated"`
		Blocks_Freed                         int32 `bson:"blocks freed"`
		Checkpoint_Size                      int32 `bson:"checkpoint size"`
		File_Allocation_Unit_Size            int32 `bson:"file allocation unit size"`
		File_Bytes_Available_For_Reuse       int32 `bson:"file bytes available for reuse"`
		File_Magic_Number                    int32 `bson:"file magic number"`
		File_Major_Version_Number            int32 `bson:"file major version number"`
		File_Size_In_Bytes                   int32 `bson:"file size in bytes"`
		Minor_Version_Number                 int32 `bson:"minor version number"`
	} `bson:"block-manager"`
	Btree struct {
		Btree_Checkpoint_Generation                   int32 `bson:"btree checkpoint generation"`
		Column_Store_Fixed_Size_Leaf_Pages            int32 `bson:"column-store fixed-size leaf pages"`
		Column_Store_Internal_Pages                   int32 `bson:"column-store internal pages"`
		Column_Store_Variable_Size_RLE_Encoded_Values int32 `bson:"column-store variable-size RLE encoded values"`
		Column_Store_Variable_Size_Deleted_Values     int32 `bson:"column-store variable-size deleted values"`
		Column_Store_Variable_Size_Leaf_Pages         int32 `bson:"column-store variable-size leaf pages"`
		Fixed_Record_Size                             int32 `bson:"fixed-record size"`
		Maximum_Internal_Page_Key_Size                int32 `bson:"maximum internal page key size"`
		Maximum_Internal_Page_Size                    int32 `bson:"maximum internal page size"`
		Maximum_Leaf_Page_Key_Size                    int32 `bson:"maximum leaf page key size"`
		Maximum_Leaf_Page_Size                        int32 `bson:"maximum leaf page size"`
		Maximum_Leaf_Page_Value_Size                  int32 `bson:"maximum leaf page value size"`
		Maximum_Tree_Depth                            int32 `bson:"maximum tree depth"`
		Number_Of_Key_Value_Pairs                     int32 `bson:"number of key/value pairs"`
		Overflow_Pages                                int32 `bson:"overflow pages"`
		Pages_Rewritten_By_Compaction                 int32 `bson:"pages rewritten by compaction"`
		Row_Store_Empty_Values                        int32 `bson:"row-store empty values"`
		Row_Store_Internal_Pages                      int32 `bson:"row-store internal pages"`
		Row_Store_Leaf_Pages                          int32 `bson:"row-store leaf pages"`
	}
	Cache struct {
		Bytes_Currently_In_The_Cache                                                        int32 `bson:"bytes currently in the cache"`
		Bytes_Dirty_In_The_Cache_Cumulative                                                 int32 `bson:"bytes dirty in the cache cumulative"`
		Bytes_Read_Into_Cache                                                               int32 `bson:"bytes read into cache"`
		Bytes_Written_From_Cache                                                            int32 `bson:"bytes written from cache"`
		Checkpoint_Blocked_Page_Eviction                                                    int32 `bson:"checkpoint blocked page eviction"`
		Data_Source_Pages_Selected_For_Eviction_Unable_To_Be_Evicted                        int32 `bson:"data source pages selected for eviction unable to be evicted"`
		Eviction_Walk_Passes_Of_A_File                                                      int32 `bson:"eviction walk passes of a file"`
		Eviction_Walk_Target_Pages_Histogram_0_9                                            int32 `bson:"eviction walk target pages histogram - 0-9"`
		Eviction_Walk_Target_Pages_Histogram_128_And_Higher                                 int32 `bson:"eviction walk target pages histogram - 128 and higher"`
		Eviction_Walk_Target_Pages_Histogram_32_63                                          int32 `bson:"eviction walk target pages histogram - 32-63"`
		Eviction_Walk_Target_Pages_Histogram_64_128                                         int32 `bson:"eviction walk target pages histogram - 64-128"`
		Eviction_Walks_Abandoned                                                            int32 `bson:"eviction walks abandoned"`
		Eviction_Walks_Gave_Up_Because_They_Restarted_Their_Walk_Twice                      int32 `bson:"eviction walks gave up because they restarted their walk twice"`
		Eviction_Walks_Gave_Up_Because_They_Saw_Too_Many_Pages_And_Found_No_Candidates      int32 `bson:"eviction walks gave up because they saw too many pages and found no candidates"`
		Eviction_Walks_Gave_Up_Because_They_Saw_Too_Many_Pages_And_Found_Too_Few_Candidates int32 `bson:"eviction walks gave up because they saw too many pages and found too few candidates"`
		Eviction_Walks_Reached_End_Of_Tree                                                  int32 `bson:"eviction walks reached end of tree"`
		Eviction_Walks_Started_From_Root_Of_Tree                                            int32 `bson:"eviction walks started from root of tree"`
		Eviction_Walks_Started_From_Saved_Location_In_Tree                                  int32 `bson:"eviction walks started from saved location in tree"`
		Hazard_Pointer_Blocked_Page_Eviction                                                int32 `bson:"hazard pointer blocked page eviction"`
		In_Memory_Page_Passed_Criteria_To_Be_Split                                          int32 `bson:"in-memory page passed criteria to be split"`
		In_Memory_Page_Splits                                                               int32 `bson:"in-memory page splits"`
		Internal_Pages_Evicted                                                              int32 `bson:"internal pages evicted"`
		Internal_Pages_Split_During_Eviction                                                int32 `bson:"internal pages split during eviction"`
		Leaf_Pages_Split_During_Eviction                                                    int32 `bson:"leaf pages split during eviction"`
		Modified_Pages_Evicted                                                              int32 `bson:"modified pages evicted"`
		Overflow_Pages_Read_Into_Cache                                                      int32 `bson:"overflow pages read into cache"`
		Page_Split_During_Eviction_Deepened_The_Tree                                        int32 `bson:"page split during eviction deepened the tree"`
		Page_Written_Requiring_Cache_Overflow_Records                                       int32 `bson:"page written requiring cache overflow records"`
		Pages_Read_Into_Cache                                                               int32 `bson:"pages read into cache"`
		Pages_Read_Into_Cache_After_Truncate                                                int32 `bson:"pages read into cache after truncate"`
		Pages_Read_Into_Cache_After_Truncate_In_Prepare_State                               int32 `bson:"pages read into cache after truncate in prepare state"`
		Pages_Read_Into_Cache_Requiring_Cache_Overflow_Entries                              int32 `bson:"pages read into cache requiring cache overflow entries"`
		Pages_Requested_From_The_Cache                                                      int32 `bson:"pages requested from the cache"`
		Pages_Seen_By_Eviction_Walk                                                         int32 `bson:"pages seen by eviction walk"`
		Pages_Written_From_Cache                                                            int32 `bson:"pages written from cache"`
		Pages_Written_Requiring_In_Memory_Restoration                                       int32 `bson:"pages written requiring in-memory restoration"`
		Tracked_Dirty_Bytes_In_The_Cache                                                    int32 `bson:"tracked dirty bytes in the cache"`
		Unmodified_Pages_Evicted                                                            int32 `bson:"unmodified pages evicted"`
	}
	Cache_walk struct {
		Average_Difference_Between_Current_Eviction_Generation_When_The_Page_Was_Last_Considered int32 `bson:"Average difference between current eviction generation when the page was last considered"`
		Average_On_Disk_Page_Image_Size_Seen                                                     int32 `bson:"Average on-disk page image size seen"`
		Average_Time_In_Cache_For_Pages_That_Have_Been_Visited_By_The_Eviction_Server            int32 `bson:"Average time in cache for pages that have been visited by the eviction server"`
		Average_Time_In_Cache_For_Pages_That_Have_Not_Been_Visited_By_The_Eviction_Server        int32 `bson:"Average time in cache for pages that have not been visited by the eviction server"`
		Clean_Pages_Currently_In_Cache                                                           int32 `bson:"Clean pages currently in cache"`
		Current_Eviction_Generation                                                              int32 `bson:"Current eviction generation"`
		Dirty_Pages_Currently_In_Cache                                                           int32 `bson:"Dirty pages currently in cache"`
		Entries_In_The_Root_Page                                                                 int32 `bson:"Entries in the root page"`
		Internal_Pages_Currently_In_Cache                                                        int32 `bson:"Internal pages currently in cache"`
		Leaf_Pages_Currently_In_Cache                                                            int32 `bson:"Leaf pages currently in cache"`
		Maximum_Difference_Between_Current_Eviction_Generation_When_The_Page_Was_Last_Considered int32 `bson:"Maximum difference between current eviction generation when the page was last considered"`
		Maximum_Page_Size_Seen                                                                   int32 `bson:"Maximum page size seen"`
		Minimum_On_Disk_Page_Image_Size_Seen                                                     int32 `bson:"Minimum on-disk page image size seen"`
		Number_Of_Pages_Never_Visited_By_Eviction_Server                                         int32 `bson:"Number of pages never visited by eviction server"`
		On_Disk_Page_Image_Sizes_Smaller_Than_A_Single_Allocation_Unit                           int32 `bson:"On-disk page image sizes smaller than a single allocation unit"`
		Pages_Created_In_Memory_And_Never_Written                                                int32 `bson:"Pages created in memory and never written"`
		Pages_Currently_Queued_For_Eviction                                                      int32 `bson:"Pages currently queued for eviction"`
		Pages_That_Could_Not_Be_Queued_For_Eviction                                              int32 `bson:"Pages that could not be queued for eviction"`
		Refs_Skipped_During_Cache_Traversal                                                      int32 `bson:"Refs skipped during cache traversal"`
		Size_Of_The_Root_Page                                                                    int32 `bson:"Size of the root page"`
		Total_Number_Of_Pages_Currently_In_Cache                                                 int32 `bson:"Total number of pages currently in cache"`
	} `bson:"cache_walk"`
	Compression struct {
		Compressed_Page_Maximum_Internal_Page_Size_Prior_To_Compression int32 `bson:"compressed page maximum internal page size prior to compression"`
		Compressed_Page_Maximum_Leaf_Page_Size_Prior_To_Compression     int32 `bson:"compressed page maximum leaf page size prior to compression"`
		Compressed_Pages_Read                                           int32 `bson:"compressed pages read"`
		Compressed_Pages_Written                                        int32 `bson:"compressed pages written"`
		Page_Written_Failed_To_Compress                                 int32 `bson:"page written failed to compress"`
		Page_Written_Was_Too_Small_To_Compress                          int32 `bson:"page written was too small to compress"`
	}
	Cursor struct {
		Bulk_Loaded_Cursor_Insert_Calls     int32 `bson:"bulk loaded cursor insert calls"`
		Cache_Cursors_Reuse_Count           int32 `bson:"cache cursors reuse count"`
		Close_Calls_That_Result_In_Cache    int32 `bson:"close calls that result in cache"`
		Create_Calls                        int32 `bson:"create calls"`
		Insert_Calls                        int32 `bson:"insert calls"`
		Insert_Key_And_Value_Bytes          int32 `bson:"insert key and value bytes"`
		Modify                              int32 `bson:"modify"`
		Modify_Key_And_Value_Bytes_Affected int32 `bson:"modify key and value bytes affected"`
		Modify_Value_Bytes_Modified         int32 `bson:"modify value bytes modified"`
		Next_Calls                          int32 `bson:"next calls"`
		Open_Cursor_Count                   int32 `bson:"open cursor count"`
		Operation_Restarted                 int32 `bson:"operation restarted"`
		Prev_Calls                          int32 `bson:"prev calls"`
		Remove_Calls                        int32 `bson:"remove calls"`
		Remove_Key_Bytes_Removed            int32 `bson:"remove key bytes removed"`
		Reserve_Calls                       int32 `bson:"reserve calls"`
		Reset_Calls                         int32 `bson:"reset calls"`
		Search_Calls                        int32 `bson:"search calls"`
		Search_Near_Calls                   int32 `bson:"search near calls"`
		Truncate_Calls                      int32 `bson:"truncate calls"`
		Update_Calls                        int32 `bson:"update calls"`
		Update_Key_And_Value_Bytes          int32 `bson:"update key and value bytes"`
		Update_Value_Size_Change            int32 `bson:"update value size change"`
	}
	Reconciliation struct {
		Dictionary_Matches                                         int32 `bson:"dictionary matches"`
		Fast_Path_Pages_Deleted                                    int32 `bson:"fast-path pages deleted"`
		Internal_Page_Key_Bytes_Discarded_Using_Suffix_Compression int32 `bson:"internal page key bytes discarded using suffix compression"`
		Internal_Page_Multi_Block_Writes                           int32 `bson:"internal page multi-block writes"`
		Internal_Page_Overflow_Keys                                int32 `bson:"internal-page overflow keys"`
		Leaf_Page_Key_Bytes_Discarded_Using_Prefix_Compression     int32 `bson:"leaf page key bytes discarded using prefix compression"`
		Leaf_Page_Multi_Block_Writes                               int32 `bson:"leaf page multi-block writes"`
		Leaf_Page_Overflow_Keys                                    int32 `bson:"leaf-page overflow keys"`
		Maximum_Blocks_Required_For_A_Page                         int32 `bson:"maximum blocks required for a page"`
		Overflow_Values_Written                                    int32 `bson:"overflow values written"`
		Page_Checksum_Matches                                      int32 `bson:"page checksum matches"`
		Page_Reconciliation_Calls                                  int32 `bson:"page reconciliation calls"`
		Page_Reconciliation_Calls_For_Eviction                     int32 `bson:"page reconciliation calls for eviction"`
		Pages_Deleted                                              int32 `bson:"pages deleted"`
	}
	session struct {
		Object_Compaction                        int32 `bson:"object compaction"`
		Tiered_Operations_Dequeued_And_Processed int32 `bson:"tiered operations dequeued and processed"`
		Tiered_Operations_Scheduled              int32 `bson:"tiered operations scheduled"`
		Tiered_Storage_Local_Retention_Time      int32 `bson:"tiered storage local retention time (secs)"`
	}
	Transaction struct {
		Race_To_Read_Prepared_Update_Retry                                                     int32 `bson:"race to read prepared update retry"`
		Rollback_To_Stable_History_Store_Records_With_Stop_Timestamps_Older_Than_Newer_Records int32 `bson:"rollback to stable history store records with stop timestamps older than newer records"`
		Rollback_To_Stable_Inconsistent_Checkpoint                                             int32 `bson:"rollback to stable inconsistent checkpoint"`
		Rollback_To_Stable_Keys_Removed                                                        int32 `bson:"rollback to stable keys removed"`
		Rollback_To_Stable_Keys_Restored                                                       int32 `bson:"rollback to stable keys restored"`
		Rollback_To_Stable_Restored_Tombstones_From_History_Store                              int32 `bson:"rollback to stable restored tombstones from history store"`
		Rollback_To_Stable_Restored_Updates_From_History_Store                                 int32 `bson:"rollback to stable restored updates from history store"`
		Rollback_To_Stable_Skipping_Delete_Rle                                                 int32 `bson:"rollback to stable skipping delete rle"`
		Rollback_To_Stable_Skipping_Stable_Rle                                                 int32 `bson:"rollback to stable skipping stable rle"`
		Rollback_To_Stable_Sweeping_History_Store_Keys                                         int32 `bson:"rollback to stable sweeping history store keys"`
		Rollback_To_Stable_Updates_Removed_From_History_Store                                  int32 `bson:"rollback to stable updates removed from history store"`
		Transaction_Checkpoints_Due_To_Obsolete_Pages                                          int32 `bson:"transaction checkpoints due to obsolete pages"`
		Update_Conflicts                                                                       int32 `bson:"update conflicts"`
	}
}
